package interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleFiveThree(t *testing.T) {
	// Create solution function that accept 1 parameter that be will multiple number 3 and 5 while each result of that multiplication is still lower than parameter inputed

	// solution (10) => 23 = 3 + 5 + 6 + 9
	// solution (20) => 78 = 3 + 5 + 6 + 9 + 10 + 12 + 15 + 18

	assert.Equal(t, multipleFiveThree(10), 23, "Multiple five three should be 23")
	assert.Equal(t, multipleFiveThree(20), 78, "Multiple five three should be 78")
}

func multipleFiveThree(num int) int {
	var result int
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}
