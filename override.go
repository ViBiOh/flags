package flags

import "strings"

type Override struct {
	value any
	name  string
}

// NewOverride replaces the default of the flags named name, ignoring case and any prefix. The value must have the Go type of the flag.
func NewOverride(name string, value any) Override {
	return Override{
		name:  name,
		value: value,
	}
}

func defaultStaticValue[T any](name string, value T, overrides []Override) T {
	for _, override := range overrides {
		if strings.EqualFold(name, override.name) {
			return override.value.(T)
		}
	}

	return value
}
