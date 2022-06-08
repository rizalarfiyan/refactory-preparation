package interview

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMiddleAlphabet(t *testing.T) {
	// Given variable string contains of all alphabet from A to Z
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Build a function that receive two paramters: first letter and last letter.
	// That will do to find middle letter between of the specified letter.
	// Example:

	// The middle between Q and U is S
	// The middle between R and U is ST
	// The middle between T and Z is W
	// The middle between Q and Z is UV

	assert.Equal(t, middleAlphabet("Q", "U"), "S", "The middle between Q and U is wrong")
	assert.Equal(t, middleAlphabet("R", "U"), "ST", "The middle between R and U is wrong")
	assert.Equal(t, middleAlphabet("T", "Z"), "W", "The middle between T and Z is wrong")
	assert.Equal(t, middleAlphabet("Q", "Z"), "UV", "The middle between Q and Z is wrong")
}

const alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func middleAlphabet(firstAlphabet, lastAlphabet string) string {
	firstAlphabet = strings.ToUpper(firstAlphabet)
	lastAlphabet = strings.ToUpper(lastAlphabet)

	//! without looping
	first := float64(strings.Index(alphabets, firstAlphabet))
	last := float64(strings.Index(alphabets, lastAlphabet))

	//! with looping
	// var first, last float64
	// for idx := range alphabets {
	// 	if first != 0 && last != 0 {
	// 		break
	// 	}
	// 	alphabet := string(alphabets[idx])
	// 	if alphabet == firstAlphabet {
	// 		first = float64(idx)
	// 	}
	// 	if alphabet == lastAlphabet {
	// 		last = float64(idx)
	// 	}
	// }

	max := math.Max(first, last)
	min := math.Min(first, last)
	center := max - min
	alpha := func(pos float64) string {
		return string(alphabets[int(pos)])
	}

	if center == 0 {
		return alpha(max)
	}

	modulo := int(center) % 2
	devide := math.Floor(center / 2)
	if modulo == 0 {
		return alpha(min + devide)
	}

	return alpha(min+devide) + alpha(min+devide+1)
}
