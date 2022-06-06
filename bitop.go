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

// func GetBitAtIndex(b uint, ind int, )

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

// RemoveBit returns the binary with the bit at index removed, length of the binary decreases by one
func RemoveBit(b uint, index int) uint {
	temp := uint(0)
	for i := 0; i < bits.Len(b); i++ {
		if i == index {
			continue
		}
		temp = temp<<1 | b>>(bits.Len(b)-i-1)&1
	}
	return temp
}

// Join returns a single binary by combining all binary values together separated by the given separator
// 0b0 is equivalent to no separator because leading zeros are omitted
func Join(bs []uint, sep uint) uint {
	joined := uint(0)
	for i, b := range bs {
		joined = joined<<bits.Len(b) | b
		if i == len(bs)-1 {
			break
		}
		joined = joined<<bits.Len(sep) | sep
	}
	return joined
}

// ColumnJoin joins the binary values in each corresponding bit position to form columns
func ColumnJoin(rows []uint, colLen int) []uint {
	cols := make([]uint, colLen)
	for i := 1; i <= colLen; i++ {
		for j := 0; j < len(rows); j++ {
			cols[i-1] = cols[i-1]<<1 | rows[j]>>(colLen-i)&1
		}
	}
	return cols
}

// Repeat returns a binary that is a repetition of the given bit pattern `b`, with length of the repeated unit `leng` (since leading zeroes would be ignored otherwise) and number of repetitions `count`
func Repeat(b uint, leng int, count int) uint {
	combined := uint(0)
	for i := 0; i < count; i++ {
		combined = combined<<leng | b
	}
	return combined
}
