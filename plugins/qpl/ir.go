package qpl

import (
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	"golang.org/x/exp/slog"

	"github.com/liam0215/anarres/runtime/plugins/qpl"
)

// The CompressionBackend IR node represents a compression service implementation
type CompressionBackend struct {
	golang.Service

	// Interfaces for generating Golang artifacts
	golang.ProvidesModule
	golang.ProvidesInterface
	golang.Instantiable

	InstanceName string
	Spec         *workflowspec.Service
}

func newCompressionBackend(name string) (*CompressionBackend, error) {
	spec, err := workflowspec.GetService[qpl.QplCompression]()
	if err != nil {
		return nil, err
	}

	node := &CompressionBackend{
		InstanceName: name,
		Spec:         spec,
	}
	return node, nil
}

// Implements ir.IRNode
func (node *CompressionBackend) Name() string {
	return node.InstanceName
}

// Implements ir.IRNode
func (node *CompressionBackend) String() string {
	return fmt.Sprintf("%v = QplCompression()", node.InstanceName)
}

// Implements golang.Service
func (node *CompressionBackend) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.Spec.Iface.ServiceInterface(ctx), nil
}

// Implements golang.ProvidesModule
func (node *CompressionBackend) AddToWorkspace(builder golang.WorkspaceBuilder) error {
	return node.Spec.AddToWorkspace(builder)
}

// Implements golang.ProvidesInterface
func (node *CompressionBackend) AddInterfaces(builder golang.ModuleBuilder) error {
	return node.Spec.AddToModule(builder)
}

// Implements golang.Instantiable
func (node *CompressionBackend) AddInstantiation(builder golang.NamespaceBuilder) error {
	if builder.Visited(node.InstanceName) {
		return nil
	}

	slog.Info(fmt.Sprintf("Instantiating QplCompression %v in %v/%v", node.InstanceName, builder.Info().Package.PackageName, builder.Info().FileName))
	return builder.DeclareConstructor(node.InstanceName, node.Spec.Constructor.AsConstructor(), nil)
}

func (node *CompressionBackend) ImplementsGolangNode()    {}
func (node *CompressionBackend) ImplementsGolangService() {}
