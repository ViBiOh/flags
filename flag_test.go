package flags_test

import (
	"flag"
	"strings"
	"testing"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	type args struct {
		defaultValue string
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      string
		wantUsage string
	}{
		"simple": {
			flags.New("name", "Name of people"),
			nil,
			args{},
			"",
			"Usage of String:\n  --name  string  Name of people ${STRING_NAME}\n",
		},
		"with default": {
			flags.New("lastname", "Last name of people"),
			nil,
			args{
				defaultValue: "Mark",
			},
			"Mark",
			"Usage of String:\n  --lastname  string  Last name of people ${STRING_LASTNAME} (default \"Mark\")\n",
		},
		"with shorthand": {
			flags.New("firstname", "First name of people").Shorthand("f"),
			nil,
			args{
				args: []string{"-f", "Mark"},
			},
			"Mark",
			"Usage of String:\n  -f, --firstname  string  First name of people ${STRING_FIRSTNAME}\n",
		},
		"with overrides": {
			flags.New("surname", "Name of people"),
			nil,
			args{
				overrides: []flags.Override{flags.NewOverride("surname", "Jane")},
			},
			"Jane",
			"Usage of String:\n  --surname  string  Name of people ${STRING_SURNAME} (default \"Jane\")\n",
		},
		"with overrides and prefix": {
			flags.New("name", "Name of people").Prefix("company"),
			nil,
			args{
				overrides: []flags.Override{flags.NewOverride("name", "ECorp")},
			},
			"ECorp",
			"Usage of String:\n  --companyName  string  [company] Name of people ${STRING_COMPANY_NAME} (default \"ECorp\")\n",
		},
		"with env": {
			flags.New("nickname", "Nickname of people"),
			func() {
				t.Setenv("STRING_NICKNAME", "John")
			},
			args{
				overrides: []flags.Override{flags.NewOverride("name", "Jane")},
			},
			"John",
			"Usage of String:\n  --nickname  string  Nickname of people ${STRING_NICKNAME} (default \"John\")\n",
		},
		"with forced env": {
			flags.New("myName", "My name").Env("MYNAME"),
			func() {
				t.Setenv("MYNAME", "John")
			},
			args{},
			"John",
			"Usage of String:\n  --myName  string  My name ${MYNAME} (default \"John\")\n",
		},
		"with args": {
			flags.New("pseudo", "Pseudo of people"),
			func() {
				t.Setenv("STRING_PSEUDO", "John")
			},
			args{
				overrides: []flags.Override{flags.NewOverride("pseudo", "Jane")},
				args:      []string{"-pseudo", "Bill"},
			},
			"Bill",
			"Usage of String:\n  --pseudo  string  Pseudo of people ${STRING_PSEUDO} (default \"John\")\n",
		},
		"with camelCase to SNAKE_CASE env variable naming": {
			flags.New("fullName", "Fullname of people"),
			func() {
				t.Setenv("STRING_FULL_NAME", "John Doe")
			},
			args{},
			"John Doe",
			"Usage of String:\n  --fullName  string  Fullname of people ${STRING_FULL_NAME} (default \"John Doe\")\n",
		},
		"with doc prefix": {
			flags.New("middlename", "Middle name of people").DocPrefix("family"),
			nil,
			args{},
			"",
			"Usage of String:\n  --middlename  string  [family] Middle name of people ${STRING_MIDDLENAME}\n",
		},
		"with prefix": {
			flags.New("name", "Name of people").Prefix("family"),
			nil,
			args{},
			"",
			"Usage of String:\n  --familyName  string  [family] Name of people ${STRING_FAMILY_NAME}\n",
		},
		"full": {
			flags.New("name", "Name of people").Prefix("full").DocPrefix("family").Env("USE_THIS_ENV"),
			func() {
				t.Setenv("USE_THIS_ENV", "John Doe")
			},
			args{
				defaultValue: "John",
				overrides:    []flags.Override{flags.NewOverride("pseudo", "Jane")},
				args:         []string{"--fullName", "Bill"},
			},
			"Bill",
			"Usage of String:\n  --fullName  string  [full] Name of people ${USE_THIS_ENV} (default \"John Doe\")\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("String", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.String(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		defaultValue int
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      int
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("studentAge", 20)},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of Int:\n  -studentA, --studentAge  int  [student] Age of people ${USE_THIS_ENV} (default 25)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Int", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Int(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		defaultValue int64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      int64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("studentAge", int64(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of Int64:\n  -studentA, --studentAge  int  [student] Age of people ${USE_THIS_ENV} (default 25)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Int64", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Int64(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		defaultValue uint
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      uint
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("studentAge", uint(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of UInt:\n  -studentA, --studentAge  uint  [student] Age of people ${USE_THIS_ENV} (default 25)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("UInt", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Uint(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestUInt64(t *testing.T) {
	type args struct {
		defaultValue uint64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      uint64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("age", uint64(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of UInt64:\n  -studentA, --studentAge  uint  [student] Age of people ${USE_THIS_ENV} (default 25)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("UInt64", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Uint64(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		defaultValue float64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      float64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "25.4")
			},
			args{
				defaultValue: 18.5,
				overrides:    []flags.Override{flags.NewOverride("studentAge", 20.3)},
				args:         []string{"--studentAge", "30.3"},
			},
			30.3,
			"Usage of Float64:\n  -studentA, --studentAge  float  [student] Age of people ${USE_THIS_ENV} (default 25.4)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Float64", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Float64(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		defaultValue bool
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      bool
		wantUsage string
	}{
		"full": {
			flags.New("active", "Control if active or not").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("a"),
			func() {
				t.Setenv("USE_THIS_ENV", "true")
			},
			args{
				defaultValue: true,
				overrides:    []flags.Override{flags.NewOverride("active", false)},
				args:         []string{"--studentActive=false"},
			},
			false,
			"Usage of Bool:\n  -studentA, --studentActive    [student] Control if active or not ${USE_THIS_ENV} (default true)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Bool", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Bool(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestDuration(t *testing.T) {
	type args struct {
		defaultValue time.Duration
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      time.Duration
		wantUsage string
	}{
		"full": {
			flags.New("remainingTime", "Remaining time in exam").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV").Shorthand("r"),
			func() {
				t.Setenv("USE_THIS_ENV", "4m30s")
			},
			args{
				defaultValue: time.Hour * 24,
				overrides:    []flags.Override{flags.NewOverride("remainingTime", time.Minute)},
				args:         []string{"--studentRemainingTime=3h"},
			},
			time.Hour * 3,
			"Usage of Duration:\n  -studentR, --studentRemainingTime  duration  [student] Remaining time in exam ${USE_THIS_ENV} (default 4m30s)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Duration", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.Duration(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}
