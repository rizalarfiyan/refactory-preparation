package interview

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternateCase(t *testing.T) {
	// Alternate each case of each of string given

	// alternateCase("abc") => "ABC"
	// alternateCase("ABC") => "abc"
	// alternateCase("Hello World") => "hELLO wORLD"

	assert.Equal(t, alternateCase("abc"), "ABC", `Alternate case "abc" should be "ABC"`)
	assert.Equal(t, alternateCase("ABC"), "abc", `Alternate case "ABC" should be "abc"`)
	assert.Equal(t, alternateCase("Hello World"), "hELLO wORLD", `Alternate case "Hello World" should be "hELLO wORLD"`)
}

func alternateCase(str string) string {
	var result string
	for _, val := range str {
		alphabet := string(val)
		if alphabet == " " {
			result += " "
			continue
		}

		if alphabet == strings.ToUpper(alphabet) {
			result += strings.ToLower(alphabet)
			continue
		}

		result += strings.ToUpper(alphabet)
	}
	return result
}
