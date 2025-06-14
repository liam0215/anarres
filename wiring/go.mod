module github.com/liam0215/anarres/wiring

go 1.23.1

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20250528164249-772aced0559e
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20250528164249-772aced0559e
	github.com/liam0215/anarres/cmplx_workload v0.0.0-00010101000000-000000000000
	github.com/liam0215/anarres/plugins v0.0.0-20250613231002-3674a1021367
	github.com/liam0215/anarres/tests v0.0.0-20250613231002-3674a1021367
	github.com/liam0215/anarres/workflow v0.0.0-20250614015759-cc451ce8844d
	github.com/liam0215/anarres/workload v0.0.0-20250613231002-3674a1021367
)

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250528164249-772aced0559e // indirect
	github.com/bradfitz/gomemcache v0.0.0-20250403215159-8d39553ac7cf // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/liam0215/anarres/runtime v0.0.0-20250613234919-63a182a37148 // indirect
	github.com/mattn/go-sqlite3 v1.14.28 // indirect
	github.com/otiai10/copy v1.14.1 // indirect
	github.com/otiai10/mint v1.6.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.17.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.36.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.36.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/exp v0.0.0-20250606033433-dcc06ee1d476 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
)

replace github.com/liam0215/anarres/workflow => ../workflow

replace github.com/liam0215/anarres/tests => ../tests

replace github.com/liam0215/anarres/workload => ../workload

replace github.com/liam0215/anarres/cmplx_workload => ../cmplx_workload

replace github.com/liam0215/anarres/plugins => ../plugins
