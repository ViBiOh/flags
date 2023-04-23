package flags_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/ViBiOh/flags"
	"github.com/stretchr/testify/assert"
)

func TestStringSlice(t *testing.T) {
	type args struct {
		defaultValue []string
		overrides    []flags.Override
		args         []string
	}

	cases := map[string]struct {
		builder   flags.Builder
		preTest   func()
		args      args
		want      []string
		wantUsage string
	}{
		"simple": {
			flags.New("tags", "Tags of ressources"),
			nil,
			args{},
			nil,
			"Usage of StringSlice:\n  --tags  string slice  Tags of ressources ${STRING_SLICE_TAGS}, as a string slice, environment variable separated by \",\"\n",
		},
		"with default value": {
			flags.New("headers", "Headers of request").Prefix("curl"),
			nil,
			args{
				defaultValue: []string{"Authorization", "Content-Type"},
			},
			[]string{"Authorization", "Content-Type"},
			"Usage of StringSlice:\n  --curlHeaders  string slice  [curl] Headers of request ${STRING_SLICE_CURL_HEADERS}, as a string slice, environment variable separated by \",\" (default [Authorization, Content-Type])\n",
		},
		"with read from environment variable": {
			flags.New("labels", "Labels of ressources").DocPrefix("metadata").EnvSeparator("|"),
			func() {
				t.Setenv("STRING_SLICE_LABELS", "env|found")
			},
			args{
				defaultValue: []string{"test", "flags"},
			},
			[]string{"env", "found"},
			"Usage of StringSlice:\n  --labels  string slice  [metadata] Labels of ressources ${STRING_SLICE_LABELS}, as a string slice, environment variable separated by \"|\" (default [env, found])\n",
		},
		"with shorthand and args": {
			flags.New("namespace", "Namespace").Shorthand("n"),
			nil,
			args{
				defaultValue: []string{"default"},
				args:         []string{"-n", "system", "--namespace", "default"},
			},
			[]string{"system", "default"},
			"Usage of StringSlice:\n  -n, --namespace  string slice  Namespace ${STRING_SLICE_NAMESPACE}, as a string slice, environment variable separated by \",\" (default [default])\n",
		},
		"with env": {
			flags.New("match", "Match").Env("MATCHES"),
			func() {
				t.Setenv("MATCHES", "info,error")
			},
			args{
				defaultValue: []string{"error"},
			},
			[]string{"info", "error"},
			"Usage of StringSlice:\n  --match  string slice  Match ${MATCHES}, as a string slice, environment variable separated by \",\" (default [info, error])\n",
		},
	}

	for intention, testCase := range cases {
		intention, testCase := intention, testCase

		t.Run(intention, func(t *testing.T) {
			fs := flag.NewFlagSet("StringSlice", flag.ContinueOnError)
			fs.Usage = flags.Usage(fs)

			var writer strings.Builder
			fs.SetOutput(&writer)

			if testCase.preTest != nil {
				testCase.preTest()
			}

			got := testCase.builder.StringSlice(fs, testCase.args.defaultValue, testCase.args.overrides)
			fs.Usage()

			assert.NoError(t, fs.Parse(testCase.args.args))
			assert.Equal(t, testCase.want, *got)
			assert.Equal(t, testCase.wantUsage, writer.String())
		})
	}
}
