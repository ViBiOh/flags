package flags

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// String creates a string flag
func String(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value string, overrides []Override) *string {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(string)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (string, error) {
		return input, nil
	})

	if len(shorthand) > 0 {
		fs.StringVar(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.StringVar(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Int creates an int flag
func Int(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value int, overrides []Override) *int {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(int)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int, error) {
		intVal, err := strconv.ParseInt(input, 10, 32)
		return int(intVal), err
	})

	if len(shorthand) > 0 {
		fs.IntVar(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.IntVar(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Int64 creates an int64 flag
func Int64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value int64, overrides []Override) *int64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(int64)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Int64Var(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Int64Var(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Uint creates an uint flag
func Uint(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value uint, overrides []Override) *uint {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(uint)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint, error) {
		intVal, err := strconv.ParseUint(input, 10, 32)
		return uint(intVal), err
	})

	if len(shorthand) > 0 {
		fs.UintVar(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.UintVar(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Uint64 creates an uint64 flag
func Uint64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value uint64, overrides []Override) *uint64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(uint64)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint64, error) {
		return strconv.ParseUint(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Uint64Var(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Uint64Var(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Float64 creates a float64 flag
func Float64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value float64, overrides []Override) *float64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(float64)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (float64, error) {
		return strconv.ParseFloat(input, 64)
	})

	if len(shorthand) > 0 {
		fs.Float64Var(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Float64Var(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Bool creates a bool flag
func Bool(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value bool, overrides []Override) *bool {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(bool)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, strconv.ParseBool)

	if len(shorthand) > 0 {
		fs.BoolVar(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.BoolVar(output, FirstLowerCase(flagName), initialValue, usage)

	return output
}

// Duration creates a duration flag
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value time.Duration, overrides []Override) *time.Duration {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	output := new(time.Duration)
	usage := formatLabel(prefix, docPrefix, label, envName)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, time.ParseDuration)

	if len(shorthand) > 0 {
		fs.DurationVar(output, FirstLowerCase(prefix+FirstUpperCase(shorthand)), initialValue, usage)
	}

	fs.DurationVar(output, FirstLowerCase(flagName), initialValue, usage)

	return output
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
func StringSlice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label string, value []string, overrides []Override) *[]string {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) ([]string, error) {
		if len(input) == 0 {
			return []string{}, nil
		}

		return strings.Split(input, ","), nil
	})

	output := new([]string)
	targetOutput := newStringSlice(initialValue, output)
	usage := formatLabel(prefix, docPrefix, label, envName) + ", as a `string slice`"

	if len(shorthand) > 0 {
		fs.Var(targetOutput, FirstLowerCase(prefix+FirstUpperCase(shorthand)), usage)
	}

	fs.Var(targetOutput, FirstLowerCase(flagName), usage)

	return output
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

func defaultValue[T any](value T, envName string, parse func(string) (T, error)) T {
	if val, ok := os.LookupEnv(envName); ok {
		parsed, err := parse(val)
		if err == nil {
			return parsed
		}
	}

	return value
}
