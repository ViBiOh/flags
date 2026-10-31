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
func String(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, value string, overrides []Override) *string {
	output := new(string)

	StringVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// StringVar bind a string flag.
func StringVar(fs *flag.FlagSet, output *string, prefix, docPrefix, name, shorthand, label, env, value string, overrides []Override) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.StringVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.StringVar(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (string, error) {
		return input, nil
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.IntVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.IntVar(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (int, error) {
		intVal, err := strconv.ParseInt(input, 10, 32)
		return int(intVal), err
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.Int64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.Int64Var(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.UintVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.UintVar(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (uint, error) {
		intVal, err := strconv.ParseUint(input, 10, 32)
		return uint(intVal), err
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.Uint64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.Uint64Var(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (uint64, error) {
		return strconv.ParseUint(input, 10, 64)
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.Float64Var(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.Float64Var(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, func(input string) (float64, error) {
		return strconv.ParseFloat(input, 64)
	})
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.BoolVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.BoolVar(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, strconv.ParseBool)
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

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		fs.DurationVar(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	fs.DurationVar(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(staticValue, envName, time.ParseDuration)
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

// Callers apply the result after registering the flag, so a value read from the environment (possibly a secret) is never shown as the default in usage.
func defaultValue[T any](value T, envName string, parse func(string) (T, error)) T {
	if val, ok := os.LookupEnv(envName); ok {
		parsed, err := parse(val)
		if err == nil {
			return parsed
		}
	}

	return value
}
