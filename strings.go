package flags

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
)

func changeFirstCase(s string, upper bool) string {
	if len(s) == 0 {
		return s
	}

	a := []rune(s)
	if upper {
		a[0] = unicode.ToUpper(a[0])
	} else {
		a[0] = unicode.ToLower(a[0])
	}

	return string(a)
}

func firstUpperCase(s string) string {
	return changeFirstCase(s, true)
}

func firstLowerCase(s string) string {
	return changeFirstCase(s, false)
}

func SnakeCase(s string) string {
	runes := []rune(s)

	var builder strings.Builder

	builder.Grow(len(s) + len(s)/4)

	for index, char := range runes {
		switch {
		case char == '-':
			builder.WriteByte('_')
		case isWordStart(runes, index):
			builder.WriteByte('_')
			builder.WriteRune(char)
		default:
			builder.WriteRune(char)
		}
	}

	return strings.ReplaceAll(builder.String(), "__", "_")
}

// An upper case letter starts a word unless it continues an acronym ("URL"),
// where only the last capital before a lower case letter does ("URLPath").
func isWordStart(runes []rune, index int) bool {
	if index == 0 || !unicode.IsUpper(runes[index]) {
		return false
	}

	previous := runes[index-1]

	switch {
	case previous == '-' || previous == '_':
		return false
	case unicode.IsUpper(previous):
		return index+1 < len(runes) && unicode.IsLower(runes[index+1])
	default:
		return true
	}
}

func Sha(content string) string {
	hasher := sha256.New()

	// no err check https://golang.org/pkg/hash/#Hash
	_, _ = fmt.Fprint(hasher, content)

	return hex.EncodeToString(hasher.Sum(nil))
}
