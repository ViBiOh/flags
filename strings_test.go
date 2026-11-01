package flags

import "testing"

func TestFirstLowerCase(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  string
	}{
		"should work with empty string": {
			"",
			"",
		},
		"should work with lower case string": {
			"test",
			"test",
		},
		"should work with regular string": {
			"OhPleaseFormatMe",
			"ohPleaseFormatMe",
		},
	}

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			if result := firstLowerCase(tc.input); result != tc.want {
				t.Errorf("FirstUpperCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestFirstUpperCase(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  string
	}{
		"should work with empty string": {
			"",
			"",
		},
		"should work with lower case string": {
			"test",
			"Test",
		},
		"should work with regular string": {
			"OhPleaseFormatMe",
			"OhPleaseFormatMe",
		},
	}

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			if result := firstUpperCase(tc.input); result != tc.want {
				t.Errorf("FirstUpperCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  string
	}{
		"should work with empty string": {
			"",
			"",
		},
		"should work with basic string": {
			"test",
			"test",
		},
		"should work with upper case starting string": {
			"OhPleaseFormatMe",
			"Oh_Please_Format_Me",
		},
		"should work with camelCase string": {
			"listCount",
			"list_Count",
		},
		"should work with dash bestween word": {
			"List-Of_thing",
			"List_Of_thing",
		},
		"should keep a trailing acronym together": {
			"httpURL",
			"http_URL",
		},
		"should split a leading acronym from the next word": {
			"URLPath",
			"URL_Path",
		},
		"should split an acronym between words": {
			"apiURLPath",
			"api_URL_Path",
		},
		"should split several acronyms": {
			"HTTPServerURL",
			"HTTP_Server_URL",
		},
		"should work with digits": {
			"Float64Slice",
			"Float64_Slice",
		},
		"should work with a dash before an upper case letter": {
			"my-CLIName",
			"my_CLI_Name",
		},
	}

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			t.Parallel()

			if result := snakeCase(tc.input); result != tc.want {
				t.Errorf("snakeCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}
