package flags

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

// String creates a string flag
func String(fs *flag.FlagSet, prefix, docPrefix, name, label string, value string, overrides []Override) *string {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.String(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (string, error) {
		return input, nil
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Int creates an int flag
func Int(fs *flag.FlagSet, prefix, docPrefix, name, label string, value int, overrides []Override) *int {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Int(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (int, error) {
		intVal, err := strconv.ParseInt(input, 10, 32)
		return int(intVal), err
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Int64 creates an int64 flag
func Int64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value int64, overrides []Override) *int64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Int64(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Uint creates an uint flag
func Uint(fs *flag.FlagSet, prefix, docPrefix, name, label string, value uint, overrides []Override) *uint {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Uint(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (uint, error) {
		intVal, err := strconv.ParseUint(input, 10, 32)
		return uint(intVal), err
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Uint64 creates an uint64 flag
func Uint64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value uint64, overrides []Override) *uint64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Uint64(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (uint64, error) {
		return strconv.ParseUint(input, 10, 64)
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Float64 creates a float64 flag
func Float64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value float64, overrides []Override) *float64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Float64(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, func(input string) (float64, error) {
		return strconv.ParseFloat(input, 64)
	}), formatLabel(prefix, docPrefix, label, envName))
}

// Bool creates a bool flag
func Bool(fs *flag.FlagSet, prefix, docPrefix, name, label string, value bool, overrides []Override) *bool {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Bool(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, strconv.ParseBool), formatLabel(prefix, docPrefix, label, envName))
}

// Duration creates a duration flag
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, label string, value time.Duration, overrides []Override) *time.Duration {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Duration(FirstLowerCase(flagName), lookupDefaultValue(name, envName, value, overrides, time.ParseDuration), formatLabel(prefix, docPrefix, label, envName))
}

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
func StringSlice(fs *flag.FlagSet, prefix, docPrefix, name, label string, values []string, overrides []Override) *[]string {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	defaultValues := lookupDefaultValue(name, envName, values, overrides, func(input string) ([]string, error) {
		if len(input) == 0 {
			return []string{}, nil
		}

		return strings.Split(input, ","), nil
	})

	p := new([]string)

	fs.Var(newStringSlice(defaultValues, p), FirstLowerCase(flagName), formatLabel(prefix, docPrefix, label, envName))

	return p
}

func getNameAndEnv(fs *flag.FlagSet, prefix, name string) (string, string) {
	name = prefix + FirstUpperCase(name)
	return name, strings.ToUpper(SnakeCase(FirstUpperCase(fs.Name()) + FirstUpperCase(name)))
}

func formatLabel(prefix, docPrefix, label, envName string) string {
	docPrefixValue := prefix
	if len(prefix) == 0 {
		docPrefixValue = docPrefix
	}

	builder := strings.Builder{}

	if len(docPrefixValue) != 0 {
		_, _ = fmt.Fprintf(&builder, "[%s] ", docPrefixValue)
	}
	_, _ = fmt.Fprintf(&builder, "%s {%s}", label, envName)

	return builder.String()
}

func lookupDefaultValue[T any](name, envName string, value T, overrides []Override, parse func(string) (T, error)) T {
	if val, ok := os.LookupEnv(envName); ok {
		parsed, err := parse(val)
		if err == nil {
			return parsed
		}
	}

	for _, override := range overrides {
		if strings.EqualFold(name, override.name) {
			return override.value.(T)
		}
	}

	return value
}
