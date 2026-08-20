module github.com/ViBiOh/flags

go 1.26.0

require github.com/stretchr/testify v1.12.1

require (
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.39.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/telemetry v0.0.0-20260811182544-a038080d80e5 // indirect
	golang.org/x/tools v0.49.1-0.20260819203639-c62e53519fb7 // indirect
	mvdan.cc/gofumpt v0.11.0 // indirect
)

tool (
	golang.org/x/tools/cmd/goimports
	golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment
	mvdan.cc/gofumpt
)
