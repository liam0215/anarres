package prometheus

import (
    "github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/address"
    "github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
    "github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
    "github.com/blueprint-uservices/blueprint/plugins/docker"
    "github.com/blueprint-uservices/blueprint/plugins/golang/goparser"
    "github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
    "github.com/blueprint-uservices/blueprint/runtime/plugins/prometheus"
)

// Blueprint IR node that represents the Prometheus container
type PrometheusCollectorContainer struct {
    docker.Container

    CollectorName string
    BindAddr      *address.BindConfig

    Iface *goparser.ParsedInterface
}

// Prometheus interface exposed to the application.
type PrometheusInterface struct {
    service.ServiceInterface
    Wrapped service.ServiceInterface
}

func (p *PrometheusInterface) GetName() string {
    return "p(" + p.Wrapped.GetName() + ")"
}

func (p *PrometheusInterface) GetMethods() []service.Method {
    return p.Wrapped.GetMethods()
}

func newPrometheusCollectorContainer(name string) (*PrometheusCollectorContainer, error) {
    spec, err := workflowspec.GetService[prometheus.PrometheusMetricCollector]()
    if err != nil {
        return nil, err
    }

    collector := &PrometheusCollectorContainer{
        CollectorName: name,
        Iface:         spec.Iface,
    }
    return collector, nil
}

func (node *PrometheusCollectorContainer) Name() string {
    return node.CollectorName
}

func (node *PrometheusCollectorContainer) String() string {
    return node.Name() + " = PrometheusCollector(" + node.BindAddr.Name() + ")"
}

func (node *PrometheusCollectorContainer) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
    iface := node.Iface.ServiceInterface(ctx)
    return &PrometheusInterface{Wrapped: iface}, nil
}

func (node *PrometheusCollectorContainer) AddContainerArtifacts(targer docker.ContainerWorkspace) error {
    return nil
}

func (node *PrometheusCollectorContainer) AddContainerInstance(target docker.ContainerWorkspace) error {
    node.BindAddr.Port = 9090    // Prometheus metrics port
    return target.DeclarePrebuiltInstance(node.CollectorName, "prom/prometheus:latest", node.BindAddr)
}
