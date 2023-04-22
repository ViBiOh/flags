package flags

import (
	"flag"
	"time"
)

type Builder struct {
	prefix       string
	docPrefix    string
	name         string
	shorthand    string
	label        string
	env          string
	envSeparator string
}

func New(name, label string) Builder {
	return Builder{
		name:         firstUpperCase(name),
		label:        label,
		envSeparator: ",",
	}
}

func (b Builder) Shorthand(shorthand string) Builder {
	b.shorthand = shorthand

	return b
}

func (b Builder) Prefix(prefix string) Builder {
	b.prefix = prefix

	return b
}

func (b Builder) DocPrefix(docPrefix string) Builder {
	b.docPrefix = docPrefix

	return b
}

func (b Builder) Env(env string) Builder {
	b.env = env

	return b
}

func (b Builder) EnvSeparator(envSeparator string) Builder {
	b.envSeparator = envSeparator

	return b
}

func (b Builder) String(fs *flag.FlagSet, value string, overrides []Override) *string {
	return String(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Int(fs *flag.FlagSet, value int, overrides []Override) *int {
	return Int(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Int64(fs *flag.FlagSet, value int64, overrides []Override) *int64 {
	return Int64(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Uint(fs *flag.FlagSet, value uint, overrides []Override) *uint {
	return Uint(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Uint64(fs *flag.FlagSet, value uint64, overrides []Override) *uint64 {
	return Uint64(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Float64(fs *flag.FlagSet, value float64, overrides []Override) *float64 {
	return Float64(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Bool(fs *flag.FlagSet, value bool, overrides []Override) *bool {
	return Bool(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Duration(fs *flag.FlagSet, value time.Duration, overrides []Override) *time.Duration {
	return Duration(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) StringSlice(fs *flag.FlagSet, value []string, overrides []Override) *[]string {
	return StringSlice(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, value, overrides)
}
