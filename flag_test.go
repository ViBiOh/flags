package flags_test

import (
	"flag"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue string
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
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
			func(t *testing.T) {
				setenv(t, "STRING_NICKNAME", "John")
			},
			args{
				overrides: []flags.Override{flags.NewOverride("nickname", "Jane")},
			},
			"John",
			"Usage of String:\n  --nickname  string  Nickname of people ${STRING_NICKNAME} (default \"Jane\")\n",
		},
		"with forced env": {
			flags.New("myName", "My name").Env("MYNAME"),
			func(t *testing.T) {
				setenv(t, "MYNAME", "John")
			},
			args{},
			"John",
			"Usage of String:\n  --myName  string  My name ${MYNAME}\n",
		},
		"with args": {
			flags.New("pseudo", "Pseudo of people"),
			func(t *testing.T) {
				setenv(t, "STRING_PSEUDO", "John")
			},
			args{
				overrides: []flags.Override{flags.NewOverride("pseudo", "Jane")},
				args:      []string{"-pseudo", "Bill"},
			},
			"Bill",
			"Usage of String:\n  --pseudo  string  Pseudo of people ${STRING_PSEUDO} (default \"Jane\")\n",
		},
		"with acronym in env variable naming": {
			flags.New("apiURLPath", "Path of the API"),
			func(t *testing.T) {
				setenv(t, "STRING_API_URL_PATH", "/v1")
			},
			args{},
			"/v1",
			"Usage of String:\n  --apiURLPath  string  Path of the API ${STRING_API_URL_PATH}\n",
		},
		"with camelCase to SNAKE_CASE env variable naming": {
			flags.New("fullName", "Fullname of people"),
			func(t *testing.T) {
				setenv(t, "STRING_FULL_NAME", "John Doe")
			},
			args{},
			"John Doe",
			"Usage of String:\n  --fullName  string  Fullname of people ${STRING_FULL_NAME}\n",
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
			flags.New("name", "Name of people").Prefix("full").DocPrefix("family").Env("USE_THIS_ENV_STRING"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_STRING", "John Doe")
			},
			args{
				defaultValue: "John",
				overrides:    []flags.Override{flags.NewOverride("name", "Jane")},
				args:         []string{"--fullName", "Bill"},
			},
			"Bill",
			"Usage of String:\n  --fullName  string  [full] Name of people ${USE_THIS_ENV_STRING} (default \"Jane\")\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("String", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue int
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      int
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_INT").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_INT", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("age", 20)},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of Int:\n  -studentA, --studentAge  int  [student] Age of people ${USE_THIS_ENV_INT} (default 20)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("Int", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestInt64(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue int64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      int64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_INT64").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_INT64", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("age", int64(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of Int64:\n  -studentA, --studentAge  int  [student] Age of people ${USE_THIS_ENV_INT64} (default 20)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("Int64", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestUInt(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue uint
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      uint
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_UINT").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_UINT", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("age", uint(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of UInt:\n  -studentA, --studentAge  uint  [student] Age of people ${USE_THIS_ENV_UINT} (default 20)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("UInt", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestUInt64(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue uint64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      uint64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_UINT64").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_UINT64", "25")
			},
			args{
				defaultValue: 18,
				overrides:    []flags.Override{flags.NewOverride("age", uint64(20))},
				args:         []string{"--studentAge", "30"},
			},
			30,
			"Usage of UInt64:\n  -studentA, --studentAge  uint  [student] Age of people ${USE_THIS_ENV_UINT64} (default 20)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("UInt64", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestFloat64(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue float64
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      float64
		wantUsage string
	}{
		"full": {
			flags.New("age", "Age of people").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_FLOAT64").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_FLOAT64", "25.4")
			},
			args{
				defaultValue: 18.5,
				overrides:    []flags.Override{flags.NewOverride("age", 20.3)},
				args:         []string{"--studentAge", "30.3"},
			},
			30.3,
			"Usage of Float64:\n  -studentA, --studentAge  float  [student] Age of people ${USE_THIS_ENV_FLOAT64} (default 20.3)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("Float64", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestBool(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue bool
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      bool
		wantUsage string
	}{
		"full": {
			flags.New("active", "Control if active or not").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_BOOL").Shorthand("a"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_BOOL", "true")
			},
			args{
				defaultValue: true,
				overrides:    []flags.Override{flags.NewOverride("active", false)},
				args:         []string{"--studentActive=false"},
			},
			false,
			"Usage of Bool:\n  -studentA, --studentActive    [student] Control if active or not ${USE_THIS_ENV_BOOL} (default false)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("Bool", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestDuration(t *testing.T) {
	t.Parallel()

	type args struct {
		defaultValue time.Duration
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func(t *testing.T)
		args      args
		want      time.Duration
		wantUsage string
	}{
		"full": {
			flags.New("remainingTime", "Remaining time in exam").Prefix("student").DocPrefix("person").Env("USE_THIS_ENV_DURATION").Shorthand("r"),
			func(t *testing.T) {
				setenv(t, "USE_THIS_ENV_DURATION", "4m30s")
			},
			args{
				defaultValue: time.Hour * 24,
				overrides:    []flags.Override{flags.NewOverride("remainingTime", time.Minute)},
				args:         []string{"--studentRemainingTime=3h"},
			},
			time.Hour * 3,
			"Usage of Duration:\n  -studentR, --studentRemainingTime  duration  [student] Remaining time in exam ${USE_THIS_ENV_DURATION} (default 1m0s)\n",
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			flagSet := flag.NewFlagSet("Duration", flag.ContinueOnError)
			flagSet.Usage = flags.Usage(flagSet)

			var writer strings.Builder
			flagSet.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			got := testCase.builder.Flag(flagSet, testCase.args.defaultValue, testCase.args.overrides)
			flagSet.Usage()

			assert.NoError(t, flagSet.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}

func TestInvalidEnvironmentValue(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		register func(flagSet *flag.FlagSet) func() any
		envName  string
		want     any
	}{
		"int": {
			func(flagSet *flag.FlagSet) func() any {
				got := flags.New("age", "Age").Flag(flagSet, 18, nil)

				return func() any { return *got }
			},
			"INVALID_AGE",
			18,
		},
		"bool": {
			func(flagSet *flag.FlagSet) func() any {
				got := flags.New("active", "Active").Flag(flagSet, true, nil)

				return func() any { return *got }
			},
			"INVALID_ACTIVE",
			true,
		},
		"duration": {
			func(flagSet *flag.FlagSet) func() any {
				got := flags.New("timeout", "Timeout").Flag(flagSet, time.Second, nil)

				return func() any { return *got }
			},
			"INVALID_TIMEOUT",
			time.Second,
		},
		"float64 slice": {
			func(flagSet *flag.FlagSet) func() any {
				got := flags.New("coords", "Coords").Flag(flagSet, []float64{1, 2}, nil)

				return func() any { return *got }
			},
			"INVALID_COORDS",
			[]float64{1, 2},
		},
	}

	for intention, testCase := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			setenv(t, testCase.envName, "not-a-value")

			message := "invalid value for environment variable " + testCase.envName

			t.Run("continue", func(t *testing.T) {
				t.Parallel()

				var writer strings.Builder

				flagSet := flag.NewFlagSet("Invalid", flag.ContinueOnError)
				flagSet.SetOutput(&writer)

				got := testCase.register(flagSet)

				assert.Equal(t, message+"\n", writer.String())
				assert.Equal(t, testCase.want, got())
			})

			t.Run("panic", func(t *testing.T) {
				t.Parallel()

				flagSet := flag.NewFlagSet("Invalid", flag.PanicOnError)
				flagSet.SetOutput(&strings.Builder{})

				assert.PanicsWithValue(t, message, func() {
					testCase.register(flagSet)
				})
			})
		})
	}
}

func TestInvalidEnvironmentValueExits(t *testing.T) {
	t.Parallel()

	if os.Getenv("FLAGS_TEST_EXIT") == "1" {
		flagSet := flag.NewFlagSet("Invalid", flag.ExitOnError)
		flags.New("age", "Age").Flag(flagSet, 18, nil)

		return
	}

	command := exec.Command(os.Args[0], "-test.run=^TestInvalidEnvironmentValueExits$")
	command.Env = append(os.Environ(), "FLAGS_TEST_EXIT=1", "INVALID_AGE=not-a-value")

	output, err := command.CombinedOutput()

	var exitErr *exec.ExitError

	assert.ErrorAs(t, err, &exitErr)
	assert.Equal(t, 2, exitErr.ExitCode())
	assert.Contains(t, string(output), "invalid value for environment variable INVALID_AGE")
	assert.NotContains(t, string(output), "not-a-value")
}

// t.Setenv forbids parallel tests because the environment is process wide. Every case uses an environment variable name no other case reads or writes, which makes plain os.Setenv safe.
func setenv(t *testing.T, key, value string) {
	t.Helper()

	previous, existed := os.LookupEnv(key)

	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("setting environment variable %s: %s", key, err)
	}

	t.Cleanup(func() {
		if existed {
			_ = os.Setenv(key, previous)
		} else {
			_ = os.Unsetenv(key)
		}
	})
}

func TestBuilderFlagVar(t *testing.T) {
	t.Parallel()

	flagSet := flag.NewFlagSet("FlagVar", flag.ContinueOnError)

	var (
		port    uint
		timeout time.Duration
		tags    []string
		coords  []float64
	)

	flags.New("port", "Port").Prefix("var").FlagVar(flagSet, &port, 8080, nil)
	flags.New("timeout", "Timeout").Prefix("var").FlagVar(flagSet, &timeout, time.Second, nil)
	flags.New("tags", "Tags").Prefix("var").FlagVar(flagSet, &tags, []string{"a"}, nil)
	flags.New("coords", "Coords").Prefix("var").FlagVar(flagSet, &coords, []float64{1}, nil)

	assert.NoError(t, flagSet.Parse([]string{"--varPort", "9090", "--varTimeout", "2s", "--varTags", "x", "--varCoords", "2.5"}))
	assert.Equal(t, uint(9090), port)
	assert.Equal(t, 2*time.Second, timeout)
	assert.Equal(t, []string{"x"}, tags)
	assert.Equal(t, []float64{2.5}, coords)
}

func TestIntegerEnvironmentValueBeyond32Bits(t *testing.T) {
	t.Parallel()

	if strconv.IntSize < 64 {
		t.Skip("int is 32 bits wide")
	}

	setenv(t, "WIDE_BIG_INT", "3000000000")
	setenv(t, "WIDE_BIG_UINT", "3000000000")

	flagSet := flag.NewFlagSet("Wide", flag.ContinueOnError)
	flagSet.SetOutput(&strings.Builder{})

	gotInt := flags.New("bigInt", "Big int").Flag(flagSet, 1, nil)
	gotUint := flags.New("bigUint", "Big uint").Flag(flagSet, uint(1), nil)

	assert.Equal(t, 3000000000, *gotInt)
	assert.Equal(t, uint(3000000000), *gotUint)
}

func TestUsageKeepsUnrelatedFlagsApart(t *testing.T) {
	t.Parallel()

	var writer strings.Builder

	flagSet := flag.NewFlagSet("Merge", flag.ContinueOnError)
	flagSet.SetOutput(&writer)
	flagSet.Usage = flags.Usage(flagSet)

	flagSet.Bool("alpha", false, "")
	flagSet.Bool("beta", false, "")
	flags.New("name", "Name").Shorthand("n").Flag(flagSet, "", nil)

	flagSet.Usage()

	assert.Equal(t, "Usage of Merge:\n      --alpha           (default false)\n      --beta            (default false)\n  -n, --name   string  Name ${MERGE_NAME}\n", writer.String())
}
