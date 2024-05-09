package flags

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// String creates a string flag.
func String(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value string, overrides []Override) *string {
	output := new(string)

	StringVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// StringVar bind a string flag.
func StringVar(fs *flag.FlagSet, output *string, prefix, docPrefix, name, shorthand, label, env string, value string, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (string, error) {
		return input, nil
	})

	if len(shorthand) > 0 {
		fs.StringVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.StringVar(output, firstLowerCase(flagName), initialValue, usage)
}

// Int creates an int flag.
func Int(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value int, overrides []Override) *int {
	output := new(int)

	IntVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// IntVar bind an int flag.
func IntVar(fs *flag.FlagSet, output *int, prefix, docPrefix, name, shorthand, label, env string, value int, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int, error) {
		intVal, err := strconv.ParseInt(input, 10, 32)
		return int(intVal), err
	})

	if len(shorthand) > 0 {
		fs.IntVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.IntVar(output, firstLowerCase(flagName), initialValue, usage)
}

// Int64 creates an int64 flag.
func Int64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value int64, overrides []Override) *int64 {
	output := new(int64)

	Int64Var(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// Int64Var bind an int64 flag.
func Int64Var(fs *flag.FlagSet, output *int64, prefix, docPrefix, name, shorthand, label, env string, value int64, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Int64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Int64Var(output, firstLowerCase(flagName), initialValue, usage)
}

// Uint creates an uint flag.
func Uint(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value uint, overrides []Override) *uint {
	output := new(uint)

	UintVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// UintVar bind an uint flag.
func UintVar(fs *flag.FlagSet, output *uint, prefix, docPrefix, name, shorthand, label, env string, value uint, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint, error) {
		intVal, err := strconv.ParseUint(input, 10, 32)
		return uint(intVal), err
	})

	if len(shorthand) > 0 {
		fs.UintVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.UintVar(output, firstLowerCase(flagName), initialValue, usage)
}

// Uint64 creates an uint64 flag.
func Uint64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value uint64, overrides []Override) *uint64 {
	output := new(uint64)

	Uint64Var(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// Uint64Var binds an uint64 flag.
func Uint64Var(fs *flag.FlagSet, output *uint64, prefix, docPrefix, name, shorthand, label, env string, value uint64, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (uint64, error) {
		return strconv.ParseUint(input, 10, 64)
	})

	if len(shorthand) > 0 {
		fs.Uint64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Uint64Var(output, firstLowerCase(flagName), initialValue, usage)
}

// Float64 creates a float64 flag.
func Float64(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value float64, overrides []Override) *float64 {
	output := new(float64)

	Float64Var(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// Float64Var binds a float64 flag.
func Float64Var(fs *flag.FlagSet, output *float64, prefix, docPrefix, name, shorthand, label, env string, value float64, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) (float64, error) {
		return strconv.ParseFloat(input, 64)
	})

	if len(shorthand) > 0 {
		fs.Float64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.Float64Var(output, firstLowerCase(flagName), initialValue, usage)
}

// Bool creates a bool flag.
func Bool(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value bool, overrides []Override) *bool {
	output := new(bool)

	BoolVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// BoolVar binds a bool flag.
func BoolVar(fs *flag.FlagSet, output *bool, prefix, docPrefix, name, shorthand, label, env string, value bool, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, strconv.ParseBool)

	if len(shorthand) > 0 {
		fs.BoolVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.BoolVar(output, firstLowerCase(flagName), initialValue, usage)
}

// Duration creates a duration flag.
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value time.Duration, overrides []Override) *time.Duration {
	output := new(time.Duration)

	DurationVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// DurationVar binds a duration flag.
func DurationVar(fs *flag.FlagSet, output *time.Duration, prefix, docPrefix, name, shorthand, label, env string, value time.Duration, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, time.ParseDuration)

	if len(shorthand) > 0 {
		fs.DurationVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), initialValue, usage)
	}

	fs.DurationVar(output, firstLowerCase(flagName), initialValue, usage)
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
