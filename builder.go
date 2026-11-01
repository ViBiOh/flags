package flags

import (
	"flag"
	"fmt"
	"time"
)

type Builder struct {
	prefix       string
	docPrefix    string
	name         string
	shorthand    string
	label        string
	env          string
	envSeparator string
}

func New(name, label string) Builder {
	return Builder{
		name:         firstUpperCase(name),
		label:        label,
		envSeparator: ",",
	}
}

func (b Builder) Shorthand(shorthand string) Builder {
	b.shorthand = shorthand

	return b
}

func (b Builder) Prefix(prefix string) Builder {
	b.prefix = prefix

	return b
}

func (b Builder) DocPrefix(docPrefix string) Builder {
	b.docPrefix = docPrefix

	return b
}

func (b Builder) Env(env string) Builder {
	b.env = env

	return b
}

func (b Builder) EnvSeparator(envSeparator string) Builder {
	b.envSeparator = envSeparator

	return b
}

type flagType interface {
	string | int | int64 | uint | uint64 | float64 | bool | time.Duration | []string | []float64
}

// Flag declares a flag of the type of value and returns its pointer. An untyped constant gives int or float64: use an explicit conversion (uint(8080)) or an instantiation (Flag[uint]) for other types.
func (b Builder) Flag[T flagType](fs *flag.FlagSet, value T, overrides []Override) *T {
	output := new(T)

	b.FlagVar(fs, output, value, overrides)

	return output
}

// FlagVar declares a flag of the type of value and binds it to output. See Flag for the typing of constants.
func (b Builder) FlagVar[T flagType](fs *flag.FlagSet, output *T, value T, overrides []Override) {
	switch typed := any(output).(type) {
	case *string:
		StringVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(string), overrides)
	case *int:
		IntVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(int), overrides)
	case *int64:
		Int64Var(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(int64), overrides)
	case *uint:
		UintVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(uint), overrides)
	case *uint64:
		Uint64Var(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(uint64), overrides)
	case *float64:
		Float64Var(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(float64), overrides)
	case *bool:
		BoolVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(bool), overrides)
	case *time.Duration:
		DurationVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, any(value).(time.Duration), overrides)
	case *[]string:
		StringSliceVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, any(value).([]string), overrides)
	case *[]float64:
		Float64SliceVar(fs, typed, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, any(value).([]float64), overrides)
	default:
		panic(fmt.Sprintf("unsupported flag type %T", output))
	}
}
