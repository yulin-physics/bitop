package bitop

import (
	"math/bits"
)

// ContainsOne returns true if the argument has at least one bit that is 0b1
func ContainsOne(b uint) bool {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 1 {
			return true
		}
	}
	return false
}

// ContainsZero returns true if the argument has at least one bit that is 0b0, excluding leading zeros on the left
func ContainsZero(b uint) bool {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 0 {
			return true
		}
	}
	return false
}

// LastIndexOfOne returns the last index of 0b1, excluding leading zeros on the left
// If no trailing ones are found, -1 is returned.
func LastIndexOfOne(b uint) int {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 1 {
			return bits.Len(b) - i - 1
		}
	}
	return -1
}

// LastIndexOfZero returns the last index of 0b0, excluding leading zeros on the left
// If no trailing zeros are found, -1 is returned.
func LastIndexOfZero(b uint) int {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 0 {
			return bits.Len(b) - i - 1
		}
	}
	return -1
}
