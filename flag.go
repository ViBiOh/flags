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
func String(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value string, overrides []Override) *string {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(string)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (string, error) {
		return input, nil
	})

	if len(shorthand) > 0 {
		fs.StringVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.StringVar(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Int creates an int flag
func Int(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value int, overrides []Override) *int {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(int)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int, error) {
		intVal, err := strconv.ParseInt(input, 10, 32)
		return int(intVal), err
	})

	if len(shorthand) > 0 {
		fs.IntVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.IntVar(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Int64 creates an int64 flag
func Int64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value int64, overrides []Override) *int64 {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(int64)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Int64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Int64Var(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Uint creates an uint flag
func Uint(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value uint, overrides []Override) *uint {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(uint)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint, error) {
		intVal, err := strconv.ParseUint(input, 10, 32)
		return uint(intVal), err
	})

	if len(shorthand) > 0 {
		fs.UintVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.UintVar(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Uint64 creates an uint64 flag
func Uint64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value uint64, overrides []Override) *uint64 {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(uint64)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint64, error) {
		return strconv.ParseUint(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Uint64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Uint64Var(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Float64 creates a float64 flag
func Float64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value float64, overrides []Override) *float64 {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(float64)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (float64, error) {
		return strconv.ParseFloat(input, 64)
	})

	if len(shorthand) > 0 {
		fs.Float64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Float64Var(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Bool creates a bool flag
func Bool(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value bool, overrides []Override) *bool {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(bool)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, strconv.ParseBool)

	if len(shorthand) > 0 {
		fs.BoolVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.BoolVar(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

// Duration creates a duration flag
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value time.Duration, overrides []Override) *time.Duration {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	output := new(time.Duration)
	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, time.ParseDuration)

	if len(shorthand) > 0 {
		fs.DurationVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.DurationVar(output, firstLowerCase(flagName), initialValue, usage)

	return output
}

func computeDescription(fs *flag.FlagSet, prefix, docPrefix, name, label, env string) (string, string, string) {
	flagName, envName := getNameAndEnv(fs, firstUpperCase(prefix), name, env)
	usage := formatLabel(prefix, docPrefix, label, envName)

	return flagName, envName, usage
}

func getNameAndEnv(fs *flag.FlagSet, prefix, name, env string) (string, string) {
	name = prefix + firstUpperCase(name)

	if len(env) == 0 {
		env = strings.ToUpper(SnakeCase(firstUpperCase(fs.Name()) + firstUpperCase(name)))
	}

	return name, env
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
	_, _ = fmt.Fprintf(&builder, "%s ${%s}", label, envName)

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
