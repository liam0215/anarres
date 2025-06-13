module github.com/liam0215/anarres/wiring

go 1.23.1

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20250528164249-772aced0559e
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20250528164249-772aced0559e
	github.com/liam0215/anarres/plugins v0.0.0-00010101000000-000000000000
	github.com/liam0215/anarres/tests v0.0.0-20250527012637-d050291189f0
	github.com/liam0215/anarres/workflow v0.0.0-20250527012637-d050291189f0
	github.com/liam0215/anarres/workload v0.0.0-20250527013547-3cf94a4ac663
)

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250528164249-772aced0559e // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/otiai10/copy v1.14.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/exp v0.0.0-20250506013437-ce4c2cf36ca6 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
)

replace github.com/liam0215/anarres/workflow => ../workflow

replace github.com/liam0215/anarres/tests => ../tests

replace github.com/liam0215/anarres/workload => ../workload

replace github.com/liam0215/anarres/plugins => ../plugins
