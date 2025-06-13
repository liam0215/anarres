module github.com/liam0215/anarres/runtime

go 1.23.1

replace github.com/liam0215/anarres/plugins => ../plugins

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250528164249-772aced0559e
	go.opentelemetry.io/otel/metric v1.36.0
)

require (
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/exp v0.0.0-20250506013437-ce4c2cf36ca6 // indirect
)
