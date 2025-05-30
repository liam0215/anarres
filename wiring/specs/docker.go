package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/clientpool"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/grpc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/opentelemetry"
	"github.com/blueprint-uservices/blueprint/plugins/retries"
	"github.com/blueprint-uservices/blueprint/plugins/simple"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/blueprint/plugins/workload"
	"github.com/blueprint-uservices/blueprint/plugins/zipkin"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/liam0215/anarres/workflow/frontend"
	"github.com/liam0215/anarres/workload/workloadgen"
)

// A wiring spec that deploys each service into its own Docker container and using gRPC to communicate between services.
//
// All RPC calls are retried up to 3 times.
// RPC clients use a client pool with 10 clients.
var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with gRPC.",
	Build:       makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	// Define the trace collector, which will be used by all services
	trace_collector := zipkin.Collector(spec, "zipkin")

	// Modifiers that will be applied to all services
	applyDockerDefaults := func(serviceName string, useHTTP ...bool) {
		// Golang-level modifiers that add functionality
		retries.AddRetries(spec, serviceName, 3)
		clientpool.Create(spec, serviceName, 10)
		opentelemetry.Instrument(spec, serviceName, trace_collector)
		if len(useHTTP) > 0 && useHTTP[0] {
			http.Deploy(spec, serviceName)
		} else {
			grpc.Deploy(spec, serviceName)
		}

		// Deploying to namespaces
		goproc.Deploy(spec, serviceName)
		linuxcontainer.Deploy(spec, serviceName)

		// Also add to tests
		gotests.Test(spec, serviceName)
	}

	compress_service := workflow.Service[compress.CompressService](spec, "compress_service")
	applyDockerDefaults(compress_service)

	cache := simple.Cache(spec, "cache")
	frontend_service := workflow.Service[frontend.Frontend](spec, "frontend", compress_service, cache)
	applyDockerDefaults(frontend_service)

	wlgen := workload.Generator[workloadgen.SimpleWorkload](spec, "wlgen", frontend_service)

	// Instantiate starting with the frontend which will trigger all other services to be instantiated
	// Also include the tests and wlgen
	return []string{frontend_service, wlgen, "gotests"}, nil
}
