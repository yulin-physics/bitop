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
// If no trailing ones are found, -1 is returned
func LastIndexOfOne(b uint) int {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 1 {
			return bits.Len(b) - i - 1
		}
	}
	return -1
}

// LastIndexOfZero returns the last index of 0b0, excluding leading zeros on the left
// If no trailing zeros are found, -1 is returned
func LastIndexOfZero(b uint) int {
	for i := 0; i < bits.Len(b); i++ {
		if (b>>i)&1 == 0 {
			return bits.Len(b) - i - 1
		}
	}
	return -1
}

// SplitAt returns the binary argument in two halves at the index specified [0, index), excluding leading zero bits
func SplitAt(b uint, index int) []uint {
	firstHalf := uint(0) | b>>(bits.Len(b)-index)
	return []uint{firstHalf, b ^ (firstHalf << (bits.Len(b) - index))}
}

// TruncateFromRight returns the binary truncated up to the index from the right, exclusive of the index
func TruncateFromRight(b uint, index int) uint {
	return b >> index
}

// ClearFromRight returns the binary bits set to zero up to the index from the right, exclusive of the index
func ClearFromRight(b uint, index int) uint {
	return b &^ (1<<index - 1)
}

// TruncateFromLeft returns binary truncated up to the index from the left, exclusive of the index
func TruncateFromLeft(b uint, index int) uint {
	return b &^ ((1<<index - 1) << (bits.Len(b) - index))
}
