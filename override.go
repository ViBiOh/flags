package flags

import "strings"

// Override is an override of default value
type Override struct {
	value any
	name  string
}

// NewOverride create a default override value
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
