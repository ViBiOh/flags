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

func (b Builder) Float64Slice(fs *flag.FlagSet, value []float64, overrides []Override) *[]float64 {
	return Float64Slice(fs, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, value, overrides)
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

func (b Builder) StringVar(fs *flag.FlagSet, output *string, value string, overrides []Override) {
	StringVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) IntVar(fs *flag.FlagSet, output *int, value int, overrides []Override) {
	IntVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Int64Var(fs *flag.FlagSet, output *int64, value int64, overrides []Override) {
	Int64Var(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) UintVar(fs *flag.FlagSet, output *uint, value uint, overrides []Override) {
	UintVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Uint64Var(fs *flag.FlagSet, output *uint64, value uint64, overrides []Override) {
	Uint64Var(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) Float64Var(fs *flag.FlagSet, output *float64, value float64, overrides []Override) {
	Float64Var(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) BoolVar(fs *flag.FlagSet, output *bool, value bool, overrides []Override) {
	BoolVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) DurationVar(fs *flag.FlagSet, output *time.Duration, value time.Duration, overrides []Override) {
	DurationVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, value, overrides)
}

func (b Builder) StringSliceVar(fs *flag.FlagSet, output *[]string, value []string, overrides []Override) {
	StringSliceVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, value, overrides)
}

func (b Builder) Float64SliceVar(fs *flag.FlagSet, output *[]float64, value []float64, overrides []Override) {
	Float64SliceVar(fs, output, b.prefix, b.docPrefix, b.name, b.shorthand, b.label, b.env, b.envSeparator, value, overrides)
}
