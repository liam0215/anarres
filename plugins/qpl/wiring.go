package qpl

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/pointer"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
)

// Service can be used by wiring specs to create a compression service instance with the specified name.
// In the compiled application, uses the [compression.QplCompression] implementation from the Blueprint runtime package
func Compression(spec wiring.WiringSpec, name string) string {
	// The nodes that we are defining
	backendName := name + ".backend"

	// Define the compression backend instance
	spec.Define(backendName, &CompressionBackend{}, func(namespace wiring.Namespace) (ir.IRNode, error) {
		return newCompressionBackend(backendName)
	})

	// Create a pointer to the backend instance
	pointer.CreatePointer[*CompressionBackend](spec, name, backendName)

	// Return the pointer; anybody who wants to access the compression instance should do so through the pointer
	return name
}
