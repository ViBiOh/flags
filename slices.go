package flags

import (
	"flag"
	"fmt"
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
func StringSlice(fs *flag.FlagSet, prefix, docPrefix, name, shorthand, label, env, envSeparator string, value []string, overrides []Override) *[]string {
	flagName, envName := getNameAndEnv(fs, firstUpperCase(prefix), name, env)
	usage := formatLabel(prefix, docPrefix, label, envName) + fmt.Sprintf(", as a `string slice`, environment variable separated by %q", envSeparator)

	initialValue := defaultValue(defaultStaticValue(name, value, overrides), envName, func(input string) ([]string, error) {
		if len(input) == 0 {
			return []string{}, nil
		}

		return strings.Split(input, envSeparator), nil
	})

	output := new([]string)
	targetOutput := newStringSlice(initialValue, output)

	if len(shorthand) > 0 {
		fs.Var(targetOutput, firstLowerCase(prefix+firstUpperCase(shorthand)), usage)
	}

	fs.Var(targetOutput, firstLowerCase(flagName), usage)

	return output
}
