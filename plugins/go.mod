module github.com/liam0215/anarres/plugins

go 1.23.1

replace github.com/liam0215/anarres/runtime => ../runtime

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20250528164249-772aced0559e
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20250528164249-772aced0559e
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f
)

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250528164249-772aced0559e // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
)
