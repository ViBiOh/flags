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
	bindScalar(fs, output, fs.StringVar, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (string, error) {
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
	bindScalar(fs, output, fs.IntVar, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (int, error) {
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
	bindScalar(fs, output, fs.Int64Var, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (int64, error) {
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
	bindScalar(fs, output, fs.UintVar, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (uint, error) {
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
	bindScalar(fs, output, fs.Uint64Var, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (uint64, error) {
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
	bindScalar(fs, output, fs.Float64Var, prefix, docPrefix, name, shorthand, label, env, value, overrides, func(input string) (float64, error) {
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
	bindScalar(fs, output, fs.BoolVar, prefix, docPrefix, name, shorthand, label, env, value, overrides, strconv.ParseBool)
}

// Duration creates a duration flag.
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env string, value time.Duration, overrides []Override) *time.Duration {
	output := new(time.Duration)

	DurationVar(fs, output, prefix, docPrefix, name, shorthand, label, env, value, overrides)

	return output
}

// DurationVar binds a duration flag.
func DurationVar(fs *flag.FlagSet, output *time.Duration, prefix, docPrefix, name, shorthand, label, env string, value time.Duration, overrides []Override) {
	bindScalar(fs, output, fs.DurationVar, prefix, docPrefix, name, shorthand, label, env, value, overrides, time.ParseDuration)
}

func bindScalar[T any](fs *flag.FlagSet, output *T, bind func(*T, string, T, string), prefix, docPrefix, name, shorthand, label, env string, value T, overrides []Override, parse func(string) (T, error)) {
	flagName, envName, usage := computeDescription(fs, prefix, docPrefix, name, label, env)

	staticValue := defaultStaticValue(name, value, overrides)

	if len(shorthand) > 0 {
		bind(output, firstLowerCase(prefix+firstUpperCase(shorthand)), staticValue, usage)
	}

	bind(output, firstLowerCase(flagName), staticValue, usage)

	*output = defaultValue(fs, staticValue, envName, parse)
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
func defaultValue[T any](fs *flag.FlagSet, value T, envName string, parse func(string) (T, error)) T {
	val, ok := os.LookupEnv(envName)
	if !ok {
		return value
	}

	parsed, err := parse(val)
	if err != nil {
		reportInvalidEnv(fs, envName)

		return value
	}

	return parsed
}

// The offending value is not printed because it may be a secret.
func reportInvalidEnv(fs *flag.FlagSet, envName string) {
	message := "invalid value for environment variable " + envName

	_, _ = fmt.Fprintln(fs.Output(), message)

	switch fs.ErrorHandling() {
	case flag.ExitOnError:
		os.Exit(2)
	case flag.PanicOnError:
		panic(message)
	case flag.ContinueOnError:
	}
}
