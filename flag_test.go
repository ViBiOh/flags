package flags

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestToString(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue string
		label        string
		overrides    []Override
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			"",
			"Test flag",
			nil,
			"Usage of ToString:\n  -test string\n    \t[cli] Test flag {TO_STRING_TEST}\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			"default",
			"Test flag",
			nil,
			"Usage of ToString:\n  -contextTest string\n    \t[context] Test flag {TO_STRING_CONTEXT_TEST} (default \"default\")\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			"default",
			"Test flag",
			nil,
			"Usage of ToString:\n  -value string\n    \t[cli] Test flag {TO_STRING_VALUE} (default \"test\")\n",
		},
		{
			"override",
			"",
			"cli",
			"overriden",
			"default",
			"Test override",
			[]Override{
				NewOverride("overriden", "override"),
			},
			"Usage of ToString:\n  -overriden string\n    \t[cli] Test override {TO_STRING_OVERRIDEN} (default \"override\")\n",
		},
	}

	os.Setenv("TO_STRING_VALUE", "test")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToString", flag.ContinueOnError)
			String(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, tc.overrides)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToString() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue int
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of ToInt:\n  -test int\n    \t[cli] Test flag {TO_INT_TEST}\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of ToInt:\n  -contextTest int\n    \t[context] Test flag {TO_INT_CONTEXT_TEST} (default 8000)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of ToInt:\n  -value int\n    \t[cli] Test flag {TO_INT_VALUE} (default 6000)\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of ToInt:\n  -invalidValue int\n    \t[cli] Test flag {TO_INT_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("TO_INT_VALUE", "6000")
	os.Setenv("TO_INT_INVALID_VALUE", "test")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToInt", flag.ContinueOnError)
			Int(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToInt() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue int64
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of ToInt64:\n  -test int\n    \t[cli] Test flag {TO_INT64_TEST}\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of ToInt64:\n  -contextTest int\n    \t[context] Test flag {TO_INT64_CONTEXT_TEST} (default 8000)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of ToInt64:\n  -value int\n    \t[cli] Test flag {TO_INT64_VALUE} (default 6000)\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of ToInt64:\n  -invalidValue int\n    \t[cli] Test flag {TO_INT64_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("TO_INT64_VALUE", "6000")
	os.Setenv("TO_INT64_INVALID_VALUE", "test")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToInt64", flag.ContinueOnError)
			Int64(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToInt64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue uint
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of ToUint:\n  -test uint\n    \t[cli] Test flag {TO_UINT_TEST}\n",
		},
		{
			"uint",
			"",
			"cli",
			"test",
			uint(10),
			"Test flag",
			"Usage of ToUint:\n  -test uint\n    \t[cli] Test flag {TO_UINT_TEST} (default 10)\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of ToUint:\n  -contextTest uint\n    \t[context] Test flag {TO_UINT_CONTEXT_TEST} (default 8000)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of ToUint:\n  -value uint\n    \t[cli] Test flag {TO_UINT_VALUE} (default 6000)\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of ToUint:\n  -invalidValue uint\n    \t[cli] Test flag {TO_UINT_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("TO_UINT_VALUE", "6000")
	os.Setenv("TO_UINT_INVALID_VALUE", "-6000")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToUint", flag.ContinueOnError)
			Uint(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToUint() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue uint64
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			0,
			"Test flag",
			"Usage of ToUint64:\n  -test uint\n    \t[cli] Test flag {TO_UINT64_TEST}\n",
		},
		{
			"uint",
			"",
			"cli",
			"test",
			10,
			"Test flag",
			"Usage of ToUint64:\n  -test uint\n    \t[cli] Test flag {TO_UINT64_TEST} (default 10)\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			8000,
			"Test flag",
			"Usage of ToUint64:\n  -contextTest uint\n    \t[context] Test flag {TO_UINT64_CONTEXT_TEST} (default 8000)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			8000,
			"Test flag",
			"Usage of ToUint64:\n  -value uint\n    \t[cli] Test flag {TO_UINT64_VALUE} (default 6000)\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			8000,
			"Test flag",
			"Usage of ToUint64:\n  -invalidValue uint\n    \t[cli] Test flag {TO_UINT64_INVALID_VALUE} (default 8000)\n",
		},
	}

	os.Setenv("TO_UINT64_VALUE", "6000")
	os.Setenv("TO_UINT64_INVALID_VALUE", "-6000")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToUint64", flag.ContinueOnError)
			Uint64(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToUint64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue float64
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			float64(0),
			"Test flag",
			"Usage of ToFloat64:\n  -test float\n    \t[cli] Test flag {TO_FLOAT64_TEST}\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			float64(12.34),
			"Test flag",
			"Usage of ToFloat64:\n  -contextTest float\n    \t[context] Test flag {TO_FLOAT64_CONTEXT_TEST} (default 12.34)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			float64(12.34),
			"Test flag",
			"Usage of ToFloat64:\n  -value float\n    \t[cli] Test flag {TO_FLOAT64_VALUE} (default 34.56)\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			float64(12.34),
			"Test flag",
			"Usage of ToFloat64:\n  -invalidValue float\n    \t[cli] Test flag {TO_FLOAT64_INVALID_VALUE} (default 12.34)\n",
		},
	}

	os.Setenv("TO_FLOAT64_VALUE", "34.56")
	os.Setenv("TO_FLOAT64_INVALID_VALUE", "12.34.56")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToFloat64", flag.ContinueOnError)
			Float64(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToFloat64() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	cases := []struct {
		intention    string
		prefix       string
		docPrefix    string
		name         string
		defaultValue bool
		label        string
		want         string
	}{
		{
			"simple",
			"",
			"cli",
			"test",
			false,
			"Test flag",
			"Usage of ToBool:\n  -test\n    \t[cli] Test flag {TO_BOOL_TEST}\n",
		},
		{
			"with prefix",
			"context",
			"cli",
			"test",
			true,
			"Test flag",
			"Usage of ToBool:\n  -contextTest\n    \t[context] Test flag {TO_BOOL_CONTEXT_TEST} (default true)\n",
		},
		{
			"env",
			"",
			"cli",
			"value",
			true,
			"Test flag",
			"Usage of ToBool:\n  -value\n    \t[cli] Test flag {TO_BOOL_VALUE}\n",
		},
		{
			"invalid env",
			"",
			"cli",
			"invalidValue",
			true,
			"Test flag",
			"Usage of ToBool:\n  -invalidValue\n    \t[cli] Test flag {TO_BOOL_INVALID_VALUE} (default true)\n",
		},
	}

	os.Setenv("TO_BOOL_VALUE", "false")
	os.Setenv("TO_BOOL_INVALID_VALUE", "test")

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			fs := flag.NewFlagSet("ToBool", flag.ContinueOnError)
			Bool(fs, tc.prefix, tc.docPrefix, tc.name, tc.label, tc.defaultValue, nil)

			var writer strings.Builder
			fs.SetOutput(&writer)
			fs.Usage()

			if result := writer.String(); result != tc.want {
				t.Errorf("ToBool() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}
