// Package specs implements wiring specs for the CompressedCache application.
//
// The wiring spec can be specified using the -w option when running wiring/main.go
package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/simple"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/liam0215/anarres/workflow/frontend"
)

// A simple wiring spec that compiles all services to a single process and therefore directly invoke each other.
// No RPC, containers, processes etc. are used.
var Basic = cmdbuilder.SpecOption{
	Name:        "basic",
	Description: "A basic single-process wiring spec with no modifiers",
	Build:       makeBasicSpec,
}

func makeBasicSpec(spec wiring.WiringSpec) ([]string, error) {
	compress_service := workflow.Service[compress.CompressService](spec, "compress_service")
	cache := simple.Cache(spec, "cache")
	frontend_service := workflow.Service[frontend.Frontend](spec, "frontend", compress_service, cache)
	return []string{compress_service, frontend_service}, nil
}
