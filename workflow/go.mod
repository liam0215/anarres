module github.com/liam0215/anarres/workflow

go 1.23.1

replace github.com/liam0215/anarres/plugins => ../plugins

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250528164249-772aced0559e
	github.com/liam0215/anarres/runtime v0.0.0-20250613234919-63a182a37148
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/otel/metric v1.36.0
)

require (
	go.mongodb.org/mongo-driver v1.17.4 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/exp v0.0.0-20250606033433-dcc06ee1d476 // indirect
)
