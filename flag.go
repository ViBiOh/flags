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
	return fs.String(FirstLowerCase(flagName), LookupEnvString(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Int creates a int flag
func Int(fs *flag.FlagSet, prefix, docPrefix, name, label string, value int, overrides []Override) *int {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Int(FirstLowerCase(flagName), LookupEnvInt(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Int64 creates a int64 flag
func Int64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value int64, overrides []Override) *int64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Int64(FirstLowerCase(flagName), LookupEnvInt64(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Uint creates a uint flag
func Uint(fs *flag.FlagSet, prefix, docPrefix, name, label string, value uint, overrides []Override) *uint {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Uint(FirstLowerCase(flagName), LookupEnvUint(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Uint64 creates a uint64 flag
func Uint64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value uint64, overrides []Override) *uint64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Uint64(FirstLowerCase(flagName), LookupEnvUint64(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Float64 creates a float64 flag
func Float64(fs *flag.FlagSet, prefix, docPrefix, name, label string, value float64, overrides []Override) *float64 {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Float64(FirstLowerCase(flagName), LookupEnvFloat64(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Bool creates a bool flag
func Bool(fs *flag.FlagSet, prefix, docPrefix, name, label string, value bool, overrides []Override) *bool {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Bool(FirstLowerCase(flagName), LookupEnvBool(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
}

// Duration creates a duration flag
func Duration(fs *flag.FlagSet, prefix, docPrefix, name, label string, value time.Duration, overrides []Override) *time.Duration {
	flagName, envName := getNameAndEnv(fs, FirstUpperCase(prefix), name)
	return fs.Duration(FirstLowerCase(flagName), LookupEnvDuration(name, envName, value, overrides), formatLabel(prefix, docPrefix, label, envName))
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
		fmt.Fprintf(&builder, "[%s] ", docPrefixValue)
	}
	fmt.Fprintf(&builder, "%s {%s}", label, envName)

	return builder.String()
}

// LookupEnvString search for given key in environment
func LookupEnvString(name, envName string, value string, overrides []Override) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvInt search for given key in environment as int
func LookupEnvInt(name, envName string, value int, overrides []Override) int {
	if val, ok := os.LookupEnv(envName); ok {
		intVal, err := strconv.ParseInt(val, 10, 32)
		if err == nil {
			return int(intVal)
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvInt64 search for given key in environment as int64
func LookupEnvInt64(name, envName string, value int64, overrides []Override) int64 {
	if val, ok := os.LookupEnv(envName); ok {
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return intVal
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvUint search for given key in environment as uint
func LookupEnvUint(name, envName string, value uint, overrides []Override) uint {
	if val, ok := os.LookupEnv(envName); ok {
		intVal, err := strconv.ParseUint(val, 10, 32)
		if err == nil {
			return uint(intVal)
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvUint64 search for given key in environment as uint64
func LookupEnvUint64(name, envName string, value uint64, overrides []Override) uint64 {
	if val, ok := os.LookupEnv(envName); ok {
		intVal, err := strconv.ParseUint(val, 10, 64)
		if err == nil {
			return intVal
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvFloat64 search for given key in environment as float64
func LookupEnvFloat64(name, envName string, value float64, overrides []Override) float64 {
	if val, ok := os.LookupEnv(envName); ok {
		floatVal, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return floatVal
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvBool search for given key in environment as bool
func LookupEnvBool(name, envName string, value bool, overrides []Override) bool {
	if val, ok := os.LookupEnv(envName); ok {
		boolBal, err := strconv.ParseBool(val)
		if err == nil {
			return boolBal
		}
	}

	return getOverridenValue(overrides, name, value)
}

// LookupEnvDuration search for given key in environment as time.Duration
func LookupEnvDuration(name, envName string, value time.Duration, overrides []Override) time.Duration {
	if val, ok := os.LookupEnv(envName); ok {
		boolBal, err := time.ParseDuration(val)
		if err == nil {
			return boolBal
		}
	}

	return getOverridenValue(overrides, name, value)
}

func getOverridenValue[T any](overrides []Override, name string, value T) T {
	for _, override := range overrides {
		if strings.EqualFold(name, override.name) {
			return override.value.(T)
		}
	}

	return value
}
