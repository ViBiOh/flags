package flags

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type sliceValue[T any] struct {
	values *[]T
	parse  func(string) (T, error)
	format func(T) string
	edited bool
}

func newSliceValue[T any](initial []T, output *[]T, parse func(string) (T, error), format func(T) string) *sliceValue[T] {
	*output = slices.Clone(initial)

	return &sliceValue[T]{values: output, parse: parse, format: format}
}

func (sv *sliceValue[T]) String() string {
	if sv == nil || sv.values == nil || len(*sv.values) == 0 {
		return ""
	}

	items := make([]string, len(*sv.values))
	for index, value := range *sv.values {
		items[index] = sv.format(value)
	}

	return "[" + strings.Join(items, ", ") + "]"
}

func (sv *sliceValue[T]) Get() any {
	return *sv.values
}

func (sv *sliceValue[T]) Set(input string) error {
	value, err := sv.parse(input)
	if err != nil {
		return err
	}

	if !sv.edited {
		sv.edited = true

		*sv.values = (*sv.values)[:0]
	}

	*sv.values = append(*sv.values, value)

	return nil
}

func StringSlice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []string, overrides []Override) *[]string {
	output := new([]string)

	StringSliceVar(fs, output, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)

	return output
}

func StringSliceVar(fs *flag.FlagSet, output *[]string, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []string, overrides []Override) {
	bindSlice(fs, output, "string slice", parseString, formatString, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)
}

func Float64Slice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []float64, overrides []Override) *[]float64 {
	output := new([]float64)

	Float64SliceVar(fs, output, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)

	return output
}

func Float64SliceVar(fs *flag.FlagSet, output *[]float64, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []float64, overrides []Override) {
	bindSlice(fs, output, "float64 slice", parseFloat64, formatFloat64, prefix, docPrefix, name, shorthand, label, env, envSeparator, values, overrides)
}

func bindSlice[T any](fs *flag.FlagSet, output *[]T, typeName string, parse func(string) (T, error), format func(T) string, prefix, docPrefix, name, shorthand, label, env, envSeparator string, values []T, overrides []Override) {
	flagName, envName := getNameAndEnv(fs, firstUpperCase(prefix), name, env)
	usage := formatLabel(prefix, docPrefix, label, envName) + fmt.Sprintf(", as a `%s`, environment variable separated by %q", typeName, envSeparator)

	target := newSliceValue(defaultStaticValue(name, values, overrides), output, parse, format)

	if len(shorthand) > 0 {
		fs.Var(target, firstLowerCase(prefix+firstUpperCase(shorthand)), usage)
	}

	fs.Var(target, firstLowerCase(flagName), usage)

	*output = defaultValue(fs, *output, envName, func(input string) ([]T, error) {
		return parseSlice(input, envSeparator, parse)
	})
}

func parseSlice[T any](input, separator string, parse func(string) (T, error)) ([]T, error) {
	if len(input) == 0 {
		return []T{}, nil
	}

	parts := strings.Split(input, separator)
	output := make([]T, 0, len(parts))

	for _, part := range parts {
		value, err := parse(part)
		if err != nil {
			return nil, fmt.Errorf("parse `%s`: %w", part, err)
		}

		output = append(output, value)
	}

	return output, nil
}

func parseString(input string) (string, error) {
	return input, nil
}

func formatString(value string) string {
	return value
}

func parseFloat64(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func formatFloat64(value float64) string {
	return fmt.Sprintf("%f", value)
}
