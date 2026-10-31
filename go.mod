module github.com/ViBiOh/flags

go 1.27.0

require github.com/stretchr/testify v1.12.1

require (
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.42.0 // indirect
	golang.org/x/sync v0.24.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/telemetry v0.0.0-20261021172232-6cd571d43fbb // indirect
	golang.org/x/tools v0.52.0 // indirect
	mvdan.cc/gofumpt v0.12.0 // indirect
)

tool (
	golang.org/x/tools/cmd/goimports
	golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment
	mvdan.cc/gofumpt
)
