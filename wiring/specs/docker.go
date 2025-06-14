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
	"github.com/blueprint-uservices/blueprint/plugins/memcached"
	// "github.com/blueprint-uservices/blueprint/plugins/simple"
	// "github.com/blueprint-uservices/blueprint/plugins/opentelemetry"
	"github.com/blueprint-uservices/blueprint/plugins/retries"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/blueprint/plugins/workload"
	// "github.com/blueprint-uservices/blueprint/plugins/jaeger"
	// "github.com/liam0215/anarres/plugins/prometheus"
	"github.com/liam0215/anarres/plugins/qpl"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/liam0215/anarres/workflow/frontend"
	"github.com/liam0215/anarres/workflow/scheduler"
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
	// metric_collector := prometheus.Collector(spec, "prometheus")
	// Define the trace collector, which will be used by all services
	// trace_collector := jaeger.Collector(spec, "jaeger")

	// Modifiers that will be applied to all services
	applyDockerDefaults := func(serviceName string, useHTTP ...bool) string {
		// Golang-level modifiers that add functionality
		retries.AddRetries(spec, serviceName, 3)
		clientpool.Create(spec, serviceName, 10)
		if len(useHTTP) > 0 && useHTTP[0] {
			http.Deploy(spec, serviceName)
		} else {
			grpc.Deploy(spec, serviceName)
		}

		// Deploying to namespaces
		goproc.Deploy(spec, serviceName)
		ctr_name := linuxcontainer.Deploy(spec, serviceName)

		// Also add to tests
		gotests.Test(spec, serviceName)

		return ctr_name
	}

	compression := qpl.Compression(spec, "qpl")
	compress_service := workflow.Service[compress.CompressService](spec, "compress_service", compression)
	// opentelemetry.Instrument(spec, compress_service, metric_collector)
	applyDockerDefaults(compress_service)

	cache := memcached.Container(spec, "cache")
	// cache := simple.Cache(spec, "cache")
	frontend_service := workflow.Service[frontend.Frontend](spec, "frontend", compress_service, cache)
	frontend_ctr := applyDockerDefaults(frontend_service)

	scheduler_service := workflow.Service[*scheduler.SchedulerServiceImpl](spec, "scheduler_service", compress_service)
	goproc.Deploy(spec, scheduler_service)
	scheduler_ctr := linuxcontainer.Deploy(spec, scheduler_service)

	wlgen := workload.Generator[workloadgen.SimpleWorkload](spec, "wlgen", frontend_service)

	// Instantiate starting with the frontend which will trigger all other services to be instantiated
	// Also include the tests and wlgen
	return []string{scheduler_ctr, frontend_ctr, wlgen, "gotests"}, nil
}
