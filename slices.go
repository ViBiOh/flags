package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type stringSlice struct {
	values *[]string
	edited bool
}

func newStringSlice(val []string, p *[]string) *stringSlice {
	*p = val

	return &stringSlice{values: p}
}

func (i *stringSlice) String() string {
	if i == nil || i.values == nil || len(*i.values) == 0 {
		return ""
	}

	return "[" + strings.Join(*i.values, ", ") + "]"
}

func (i *stringSlice) Get() any {
	return *i.values
}

func (i *stringSlice) Set(value string) error {
	if !i.edited {
		i.edited = true

		*i.values = (*i.values)[:0]
	}

	*i.values = append(*i.values, value)

	return nil
}

// StringSlice creates a string slice flag
func StringSlice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []string, overrides []Override) *[]string {
	output := new([]string)

	StringSliceVar(fs, output, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)

	return output
}

// StringSliceVar binds a string slice flag
func StringSliceVar(fs *flag.FlagSet, output *[]string, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []string, overrides []Override) {
	flagName, envName := getNameAndEnv(fs, firstUpperCase(prefix), name, env)
	usage := formatLabel(prefix, docPrefix, label, envName) + fmt.Sprintf(", as a `string slice`, environment variable separated by %q", envSeparator)

	initialValue := defaultValue(defaultStaticValue(name, values, overrides), envName, func(input string) ([]string, error) {
		if len(input) == 0 {
			return []string{}, nil
		}

		return strings.Split(input, envSeparator), nil
	})

	targetOutput := newStringSlice(initialValue, output)

	if len(shorthand) > 0 {
		fs.Var(targetOutput, firstLowerCase(prefix+firstUpperCase(shorthand)), usage)
	}

	fs.Var(targetOutput, firstLowerCase(flagName), usage)
}

type float64Slice struct {
	values *[]float64
	edited bool
}

func newfloat64Slice(val []float64, p *[]float64) *float64Slice {
	*p = val

	return &float64Slice{values: p}
}

func (i *float64Slice) String() string {
	if i == nil || i.values == nil || len(*i.values) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, value := range *i.values {
		if builder.Len() != 0 {
			_, _ = fmt.Fprintf(&builder, ", ")
		}

		_, _ = fmt.Fprintf(&builder, "%f", value)
	}

	return "[" + builder.String() + "]"
}

func (i *float64Slice) Get() any {
	return *i.values
}

func (i *float64Slice) Set(value string) error {
	if !i.edited {
		i.edited = true

		*i.values = (*i.values)[:0]
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	*i.values = append(*i.values, floatValue)

	return nil
}

// Float64Slice creates a string slice flag
func Float64Slice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []float64, overrides []Override) *[]float64 {
	output := new([]float64)

	Float64SliceVar(fs, output, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)

	return output
}

// Float64SliceVar binds a string slice flag
func Float64SliceVar(fs *flag.FlagSet, output *[]float64, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []float64, overrides []Override) {
	flagName, envName := getNameAndEnv(fs, firstUpperCase(prefix), name, env)
	usage := formatLabel(prefix, docPrefix, label, envName) + fmt.Sprintf(", as a `float64 slice`, environment variable separated by %q", envSeparator)

	initialValue := defaultValue(defaultStaticValue(name, values, overrides), envName, func(input string) ([]float64, error) {
		if len(input) == 0 {
			return []float64{}, nil
		}

		parts := strings.Split(input, envSeparator)

		var output []float64

		for _, value := range parts {
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("parse `%s`: %w", value, err)
			}

			output = append(output, floatValue)
		}

		return output, nil
	})

	targetOutput := newfloat64Slice(initialValue, output)

	if len(shorthand) > 0 {
		fs.Var(targetOutput, firstLowerCase(prefix+firstUpperCase(shorthand)), usage)
	}

	fs.Var(targetOutput, firstLowerCase(flagName), usage)
}
