module github.com/liam0215/anarres/tests

go 1.24.3

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250525092121-5eadb0c0735e
	github.com/liam0215/anarres/workflow v0.0.0-20250526192138-4be3cbc1f8a4
	github.com/stretchr/testify v1.10.0
)

replace github.com/liam0215/anarres/workflow => ../workflow

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20250506013437-ce4c2cf36ca6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
