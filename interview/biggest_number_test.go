package interview

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiggestNumber(t *testing.T) {
	// Create a function that takes one positive three digit integer and rearranges its digits to get the maximum possible number. Assume that the argument is an integer.

	// Returm null if the argument is invalid.

	// maxRedigit(123) --> 321
	// maxRedigit(231) --> 321
	// maxRedigit(321) --> 321

	// maxRedigit(-1) --> null
	// maxRedigit(0) --> null
	// maxRedigit(99) --> null
	// maxRedigit(1000) --> null

	result := 321
	assert.Equal(t, maxRedigit(231), &result, "Max redigit of 231 should be 321")
	assert.Equal(t, maxRedigit(123), &result, "Max redigit of 123 should be 321")
	assert.Equal(t, maxRedigit(321), &result, "Max redigit of 321 should be 321")
	assert.Nil(t, maxRedigit(-1), "Max redigit of -1 should be nil")
	assert.Nil(t, maxRedigit(0), "Max redigit of 0 should be nil")
	assert.Nil(t, maxRedigit(99), "Max redigit of 99 should be nil")
	assert.Nil(t, maxRedigit(1000), "Max redigit of 1000 should be nil")
}

func maxRedigit(num int) *int {
	if num < 100 || num > 999 {
		return nil
	}

	var arrStr []string
	for num != 0 {
		arrStr = append(arrStr, fmt.Sprint(num%10))
		num /= 10
	}

	sort.Slice(arrStr, func(i, j int) bool {
		return arrStr[i] > arrStr[j]
	})

	strJoin := strings.Join(arrStr, "")
	res, _ := strconv.Atoi(strJoin)
	return &res
}
