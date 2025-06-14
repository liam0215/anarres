// An application for compiling the CompressedCache application.
// Provides a number of different wiring specs for compiling
// the application in different configurations.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	_ "github.com/liam0215/anarres/tests"
	"github.com/liam0215/anarres/wiring/specs"
)

func main() {
	// Make sure tests and workflow can be found
	workflowspec.AddModule("github.com/liam0215/anarres/tests")

	// Build a supported wiring spec
	name := "CompressedCache"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Basic,
		specs.GRPC,
		specs.Docker,
		specs.DockerMonolith,
	)
}
