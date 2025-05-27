// Package specs implements wiring specs for the CompressedCache application.
//
// The wiring spec can be specified using the -w option when running wiring/main.go
package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/clientpool"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/grpc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/retries"
	"github.com/blueprint-uservices/blueprint/plugins/simple"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/blueprint/plugins/workload"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/liam0215/anarres/workflow/frontend"
	"github.com/liam0215/anarres/workload/workloadgen"
)

// A simple wiring spec that compiles all services to a single process and therefore directly invoke each other.
// No RPC, containers, processes etc. are used.
var GRPC = cmdbuilder.SpecOption{
	Name:        "grpc",
	Description: "Deploys each service in a separate process with gRPC.",
	Build:       makeGrpcSpec,
}

func makeGrpcSpec(spec wiring.WiringSpec) ([]string, error) {
	// Modifiers that will be applied to all services
	applyDefaults := func(serviceName string, useHTTP ...bool) {
		// Golang-level modifiers that add functionality
		retries.AddRetries(spec, serviceName, 3)
		clientpool.Create(spec, serviceName, 3)
		if len(useHTTP) > 0 && useHTTP[0] {
			http.Deploy(spec, serviceName)
		} else {
			grpc.Deploy(spec, serviceName)
		}

		// Deploying to namespaces
		goproc.Deploy(spec, serviceName)

		// Also add to tests
		gotests.Test(spec, serviceName)
	}

	compress_service := workflow.Service[compress.CompressService](spec, "compress_service")
	applyDefaults(compress_service)

	cache := simple.Cache(spec, "cache")
	frontend_service := workflow.Service[frontend.Frontend](spec, "frontend", compress_service, cache)
	applyDefaults(frontend_service)

	wlgen := workload.Generator[workloadgen.SimpleWorkload](spec, "wlgen", frontend_service)

	// Instantiate starting with the frontend which will trigger all other services to be instantiated
	// Also include the tests and wlgen
	return []string{frontend_service, wlgen, "gotests"}, nil
}
