package qpl

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/pointer"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/liam0215/anarres/runtime/plugins/qpl"
	"github.com/liam0215/anarres/runtime/plugins/sw_qpl"
)

func HwCompression(spec wiring.WiringSpec, name string) string {
	return Compression[qpl.QplCompression](spec, name)
}

func SwCompression(spec wiring.WiringSpec, name string) string {
	return Compression[sw_qpl.SwQplCompression](spec, name)
}

// Service can be used by wiring specs to create a compression service instance with the specified name.
// In the compiled application, uses the [compression.QplCompression] implementation from the Blueprint runtime package
func Compression[BackendImpl any](spec wiring.WiringSpec, name string) string {
	// The nodes that we are defining
	backendName := name + ".backend"

	// Define the compression backend instance
	spec.Define(backendName, &CompressionBackend{}, func(namespace wiring.Namespace) (ir.IRNode, error) {
		return newCompressionBackend[BackendImpl](backendName)
	})

	// Create a pointer to the backend instance
	pointer.CreatePointer[*CompressionBackend](spec, name, backendName)

	// Return the pointer; anybody who wants to access the compression instance should do so through the pointer
	return name
}
