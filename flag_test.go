package flags

import (
	"flag"
	"os"
	"strings"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue string
		label        string
		overrides    []Override
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			"",
			"Test flag",
			nil,
			"Usage of String:\n  -test string\n    \t[cli] Test flag {STRING_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			"default",
			"Test flag",
			nil,
			"Usage of String:\n  -contextTest string\n    \t[context] Test flag {STRING_CONTEXT_TEST} (default \"default\")\n",
		},
		"env": {
			"",
			"cli",
			"value",
			"default",
			"Test flag",
			nil,
			"Usage of String:\n  -value string\n    \t[cli] Test flag {STRING_VALUE} (default \"test\")\n",
		},
		"override": {
			"",
			"cli",
			"overriden",
			"default",
			"Test override",
			[]Override{
				NewOverride("overriden", "override"),
			},
			"Usage of String:\n  -overriden string\n    \t[cli] Test override {STRING_OVERRIDEN} (default \"override\")\n",
		},
	}

	os.Setenv("STRING_VALUE", "test")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("String", flag.ContinueOnError)
			String(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, tc.overrides)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("String() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue int
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of Int:\n  -test int\n    \t[cli] Test flag {INT_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of Int:\n  -contextTest int\n    \t[context] Test flag {INT_CONTEXT_TEST} (default 8000)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of Int:\n  -value int\n    \t[cli] Test flag {INT_VALUE} (default 6000)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of Int:\n  -invalidValue int\n    \t[cli] Test flag {INT_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("INT_VALUE", "6000")
	os.Setenv("INT_INVALID_VALUE", "test")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Int", flag.ContinueOnError)
			Int(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Int() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue int64
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of Int64:\n  -test int\n    \t[cli] Test flag {INT64_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of Int64:\n  -contextTest int\n    \t[context] Test flag {INT64_CONTEXT_TEST} (default 8000)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of Int64:\n  -value int\n    \t[cli] Test flag {INT64_VALUE} (default 6000)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of Int64:\n  -invalidValue int\n    \t[cli] Test flag {INT64_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("INT64_VALUE", "6000")
	os.Setenv("INT64_INVALID_VALUE", "test")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Int64", flag.ContinueOnError)
			Int64(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Int64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue uint
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of Uint:\n  -test uint\n    \t[cli] Test flag {UINT_TEST}\n",
		},
		"uint": {
			"",
			"cli",
			"test",
			uint(10),
			"Test flag",
			"Usage of Uint:\n  -test uint\n    \t[cli] Test flag {UINT_TEST} (default 10)\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of Uint:\n  -contextTest uint\n    \t[context] Test flag {UINT_CONTEXT_TEST} (default 8000)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of Uint:\n  -value uint\n    \t[cli] Test flag {UINT_VALUE} (default 6000)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of Uint:\n  -invalidValue uint\n    \t[cli] Test flag {UINT_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("UINT_VALUE", "6000")
	os.Setenv("UINT_INVALID_VALUE", "-6000")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Uint", flag.ContinueOnError)
			Uint(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Uint() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue uint64
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of Uint64:\n  -test uint\n    \t[cli] Test flag {UINT64_TEST}\n",
		},
		"uint": {
			"",
			"cli",
			"test",
			10,
			"Test flag",
			"Usage of Uint64:\n  -test uint\n    \t[cli] Test flag {UINT64_TEST} (default 10)\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of Uint64:\n  -contextTest uint\n    \t[context] Test flag {UINT64_CONTEXT_TEST} (default 8000)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of Uint64:\n  -value uint\n    \t[cli] Test flag {UINT64_VALUE} (default 6000)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of Uint64:\n  -invalidValue uint\n    \t[cli] Test flag {UINT64_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("UINT64_VALUE", "6000")
	os.Setenv("UINT64_INVALID_VALUE", "-6000")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Uint64", flag.ContinueOnError)
			Uint64(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Uint64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue float64
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			float64(0),
			"Test flag",
			"Usage of Float64:\n  -test float\n    \t[cli] Test flag {FLOAT64_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			12.34,
			"Test flag",
			"Usage of Float64:\n  -contextTest float\n    \t[context] Test flag {FLOAT64_CONTEXT_TEST} (default 12.34)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			12.34,
			"Test flag",
			"Usage of Float64:\n  -value float\n    \t[cli] Test flag {FLOAT64_VALUE} (default 34.56)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			12.34,
			"Test flag",
			"Usage of Float64:\n  -invalidValue float\n    \t[cli] Test flag {FLOAT64_INVALID_VALUE} (default 12.34)\n",
		},
	}

	os.Setenv("FLOAT64_VALUE", "34.56")
	os.Setenv("FLOAT64_INVALID_VALUE", "12.34.56")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Float64", flag.ContinueOnError)
			Float64(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Float64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue bool
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			false,
			"Test flag",
			"Usage of Bool:\n  -test\n    \t[cli] Test flag {BOOL_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			true,
			"Test flag",
			"Usage of Bool:\n  -contextTest\n    \t[context] Test flag {BOOL_CONTEXT_TEST} (default true)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			true,
			"Test flag",
			"Usage of Bool:\n  -value\n    \t[cli] Test flag {BOOL_VALUE}\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			true,
			"Test flag",
			"Usage of Bool:\n  -invalidValue\n    \t[cli] Test flag {BOOL_INVALID_VALUE} (default true)\n",
		},
	}

	os.Setenv("BOOL_VALUE", "false")
	os.Setenv("BOOL_INVALID_VALUE", "test")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Bool", flag.ContinueOnError)
			Bool(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Bool() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue time.Duration
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of Duration:\n  -test duration\n    \t[cli] Test flag {DURATION_TEST}\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			time.Minute,
			"Test flag",
			"Usage of Duration:\n  -contextTest duration\n    \t[context] Test flag {DURATION_CONTEXT_TEST} (default 1m0s)\n",
		},
		"env": {
			"",
			"cli",
			"value",
			time.Minute,
			"Test flag",
			"Usage of Duration:\n  -value duration\n    \t[cli] Test flag {DURATION_VALUE} (default 1s)\n",
		},
		"invalid env": {
			"",
			"cli",
			"invalidValue",
			time.Minute,
			"Test flag",
			"Usage of Duration:\n  -invalidValue duration\n    \t[cli] Test flag {DURATION_INVALID_VALUE} (default 1m0s)\n",
		},
	}

	os.Setenv("DURATION_VALUE", "1s")
	os.Setenv("DURATION_INVALID_VALUE", "test")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Duration", flag.ContinueOnError)
			Duration(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Duration() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestStringSlice(t *testing.T) {
	cases := map[string]struct {
		prefix       string
		docPrefix    string
		name         string
		defaultValue []string
		label        string
		want         string
	}{
		"simple": {
			"",
			"cli",
			"test",
			nil,
			"Test flag",
			"Usage of Values:\n  -test string slice\n    \t[cli] Test flag {VALUES_TEST}, as a string slice\n",
		},
		"with prefix": {
			"context",
			"cli",
			"test",
			[]string{"value"},
			"Test flag",
			"Usage of Values:\n  -contextTest string slice\n    \t[context] Test flag {VALUES_CONTEXT_TEST}, as a string slice (default [value])\n",
		},
		"env": {
			"",
			"cli",
			"value",
			[]string{"value"},
			"Test flag",
			"Usage of Values:\n  -value string slice\n    \t[cli] Test flag {VALUES_VALUE}, as a string slice (default [overriden])\n",
		},
	}

	os.Setenv("VALUES_VALUE", "overriden")

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("Values", flag.ContinueOnError)
			StringSlice(fs, tc.prefix, tc.docPrefix, tc.name, "", tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("Duration() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}
