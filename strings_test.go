package flags

import "testing"

func TestFirstLowerCase(t *testing.T) {
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
			if result := firstLowerCase(tc.input); result != tc.want {
				t.Errorf("FirstUpperCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestFirstUpperCase(t *testing.T) {
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
			if result := firstUpperCase(tc.input); result != tc.want {
				t.Errorf("FirstUpperCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
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
	}

	for intention, tc := range cases {
		t.Run(intention, func(t *testing.T) {
			if result := SnakeCase(tc.input); result != tc.want {
				t.Errorf("SnakeCase() = `%s`, want `%s`", result, tc.want)
			}
		})
	}
}
