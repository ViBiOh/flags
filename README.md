# flags

[![Build](https://github.com/ViBiOh/flags/workflows/Build/badge.svg)](https://github.com/ViBiOh/flags/actions)
[![codecov](https://codecov.io/gh/ViBiOh/flags/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/flags)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_flags&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_flags)

Golang flags parser with zero dependency.

### Usage

See [simple.go](cmd/simple/simple.go) for basic usage.

## Concept

`flags` gives a simple way to get flag's value from argument or environment variable.

Argument takes priority over environment variable. Argument and environment variable names are generated from configuration you pass.

The [`FlagSet`](https://pkg.go.dev/flag#FlagSet) name is used to prefix all environment variable, replacing `-` by `_`.

The `prefix` name is used then to specialize your flag name (e.g. if you have want to use the same `flags` twice, you can change the prefix)

The argument's name is in camelCase format. The environement variable name is in SNAKE_UPPER_CASE format.

Flags can take a default value, that can be overriden programatically, always in the case you reuse the same `flags` twice (see [advanced.go](cmd/advanced/advanced.go) example.)

### Security

Be careful when using the arguments values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.
