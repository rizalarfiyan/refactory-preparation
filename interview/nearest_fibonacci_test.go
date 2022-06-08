package interview

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestFibonacci(t *testing.T) {
	// Have the function to find nearest Fibonacci numbers, that the parameter is array.
	// Example Input:
	// arr = [15,1,3]
	// then your program expected to made output = 2

	// Why Output is 2?

	// Because by the above input example [15,1,3], if we add up them it will 15+1+3 = 19, and the nearest fibonacci of 19 is 21, so we need to add 2 then it can be 21. So, the correct answer is 2.

	assert.Equal(t, nearestFibonacci([]int{15, 1, 3}), 2, "Nearest fibonacci should be 2")
}

func nearestFibonacci(nums []int) int {
	var num int
	for _, val := range nums {
		num += val
	}

	if num == 0 {
		return 0
	}

	first := 0
	second := 1
	third := first + second
	for third <= num {
		first, second = second, third
		third = first + second
	}

	subtraction := func(num1, num2 int) int {
		if num1 > num2 {
			return num1 - num2
		}
		return num2 - num1
	}

	if math.Abs(float64(third-num)) >= math.Abs(float64(second-num)) {
		return subtraction(num, second)
	}

	return subtraction(num, third)
}
