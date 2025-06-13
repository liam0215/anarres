package prometheus

import (
    "fmt"

    "github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/address"
    "github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
    "github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
    "github.com/blueprint-uservices/blueprint/plugins/golang"
    "github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
    "github.com/blueprint-uservices/blueprint/runtime/plugins/prometheus"
    "golang.org/x/exp/slog"
)

// Blueprint IR node representing a client to the prometheus container
type PrometheusCollectorClient struct {
    golang.Node
    service.ServiceNode
    golang.Instantiable
    ClientName string
    ServerDial *address.DialConfig

    InstanceName string
    Spec         *workflowspec.Service
}

func newPrometheusCollectorClient(name string, addr *address.DialConfig) (*PrometheusCollectorClient, error) {
    spec, err := workflowspec.GetService[prometheus.PrometheusMetricCollector]()
    if err != nil {
        return nil, err
    }

    node := &PrometheusCollectorClient{
        InstanceName: name,
        ClientName:   name,
        ServerDial:   addr,
        Spec:         spec,
    }
    return node, nil
}

// Implements ir.IRNode
func (node *PrometheusCollectorClient) Name() string {
    return node.ClientName
}

// Implements ir.IRNode
func (node *PrometheusCollectorClient) String() string {
    return node.Name() + " = PrometheusClient(" + node.ServerDial.Name() + ")"
}

// Implements golang.Instantiable
func (node *PrometheusCollectorClient) AddInstantiation(builder golang.NamespaceBuilder) error {
    // Only generate instantiation code for this instance once
    if builder.Visited(node.ClientName) {
        return nil
    }

    slog.Info(fmt.Sprintf("Instantiating PrometheusClient %v in %v/%v", node.InstanceName, builder.Info().Package.PackageName, builder.Info().FileName))

    return builder.DeclareConstructor(node.InstanceName, node.Spec.Constructor.AsConstructor(), []ir.IRNode{node.ServerDial})
}

// Implements service.ServiceNode
func (node *PrometheusCollectorClient) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
    return node.Spec.Iface.ServiceInterface(ctx), nil
}

// Implements golang.ProvidesInterface
func (node *PrometheusCollectorClient) AddInterfaces(builder golang.ModuleBuilder) error {
    return node.Spec.AddToModule(builder)
}

// Implements golang.ProvidesModule
func (node *PrometheusCollectorClient) AddToWorkspace(builder golang.WorkspaceBuilder) error {
    return node.Spec.AddToWorkspace(builder)
}

func (node *PrometheusCollectorClient) ImplementsGolangNode() {}

func (node *PrometheusCollectorClient) ImplementsOTCollectorClient() {}
