package bitop

import (
	"math/bits"
)

// Unit consists of a binary value which the bitop functions act on, and the number of bits in the binary as leng
type Unit struct {
	value uint
	leng  int
}

// NewUnit returns a new unit type required by all functions
// Give -ve leng to take default length measure of the input binary, which omits leading zeros
func NewUnit(b uint, leng int) Unit {
	if leng < 0 {
		leng = bits.Len(b)
	}
	return Unit{value: b, leng: leng}
}

// Contains returns true if the binary has at least one bit that is 0b1
func Contains(b, sub Unit) bool {
	for i := 0; i <= b.leng-sub.leng; i++ {
		window := TruncateFromLeft(b, i)
		window = TruncateFromRight(window, b.leng-i-sub.leng)
		if window == sub.value {
			return true
		}
	}
	return false
}

// LastIndex returns the last index of the given bit pattern, if no matching found -1 is returned
func LastIndex(b, sub Unit) int {
	for i := 0; i <= b.leng-sub.leng; i++ {
		window := TruncateFromLeft(b, b.leng-i-sub.leng)
		window = TruncateFromRight(window, i)
		if window == sub.value {
			return b.leng - i - sub.leng
		}
	}
	return -1
}

// GetBitAtIndex returns the bit at index `ind` of the given binary, index counting from left to right from zero as usual
func GetBitAtIndex(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value >> uint(b.leng-ind-1) & 1
}

// SplitAt returns the binary argument in two halves at the index specified [0, index)
func SplitAt(b Unit, ind int) []uint {
	if ind < 0 {
		return []uint{b.value}
	}
	firstHalf := uint(0) | b.value>>(b.leng-ind)
	return []uint{firstHalf, b.value ^ (firstHalf << (b.leng - ind))}
}

// TruncateFromRight returns the binary truncated up to the index from the right, exclusive of the index
func TruncateFromRight(b uint, pos int) uint {
	if pos < 0 {
		return b
	}
	return b >> pos
}

// ClearFromRight returns the binary bits set to zero up to the index from the right, exclusive of the index
func ClearFromRight(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value &^ (1<<ind - 1)
}

// TruncateFromLeft returns binary truncated up to the index from the left, exclusive of the index
func TruncateFromLeft(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value &^ ((1<<ind - 1) << (b.leng - ind))
}

// RemoveBit returns the binary with the bit at index removed, length of the binary decreases by one
func RemoveBit(b Unit, index int) uint {
	new := uint(0)
	for i := 0; i < b.leng; i++ {
		if i == index {
			continue
		}
		new = new<<1 | b.value>>(b.leng-i-1)&1
	}
	return new
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

// Repeat returns a binary that is a repetition of the given bit pattern `b`, with length of the repeated unit `leng` (since leading zeroes would be ignored otherwise, parse -1 for leng if length unknown) and number of repetitions `count`
func Repeat(b uint, count int, leng int) uint {
	if leng < 0 {
		leng = bits.Len(b)
	}
	combined := uint(0)
	for i := 0; i < count; i++ {
		combined = combined<<leng | b
	}
	return combined
}

// Replace returns a binary with any old bit pattern replaced by new, up to n times of occurrences
// leng, oldLeng, newLeng are optional properties for b, old, new, to account for leading zeroes
func Replace(b uint, old uint, new uint, n int, leng int, oldLeng int, newLeng int) uint {
	if leng < 0 {
		leng = bits.Len(b)
	}
	if oldLeng < 0 {
		oldLeng = bits.Len(old)
	}
	if newLeng < 0 {
		newLeng = bits.Len(new)
	}
	result := uint(0)
	for i := 0; i < leng; {
		// window := TruncateFromLeft(b, i, leng)
		// window = TruncateFromRight(window, leng-i-oldLeng)
		// if window == old && n > 0 {
		// 	result = result<<newLeng | new
		// 	n--
		// 	i += oldLeng
		// } else {
		// 	result = result<<1 | GetBitAtIndex(b, i, leng)
		// 	i++
		// }
	}
	return result
}

// FlipAtIndex flips the bit at index i in the binary
// leng is an optional input for number of bits in the binary, to account for leading zeroes (set as -1 to omit)
func FlipAtIndex(b uint, index int, leng int) uint {
	if index < 0 {
		return b
	}
	if leng < 0 {
		leng = bits.Len(b)
	}
	return b | 1<<uint(leng-index-1)
}

// Flip returns a binary with all bits flipped
// leng is an optional input for number of bits in the binary, to account for leading zeroes (set as -1 to omit)
func Flip(b uint, leng int) uint {
	if leng < 0 {
		leng = bits.Len(b)
	}
	result := uint(0)
	for i := leng - 1; i >= 0; i-- {
		result = result<<1 | (1 ^ b>>i&1)
	}
	return result
}

// Reverse returns a binary with bits in reversed order
// leng is an optional input for number of bits in the binary, to account for leading zeroes (set as -1 to omit)
func Reverse(b uint, leng int) uint {
	if leng < 0 {
		leng = bits.Len(b)
	}
	reversed := uint(0)
	for i := 0; i < leng; i++ {
		reversed = reversed<<1 | b>>i&1
	}
	return reversed
}

// IsPalindrome checks if the binary is symmetrical
// leng is an optional input for number of bits in the binary (set as -1 to omit)
func IsPalindrome(b uint, leng int) bool {
	if leng < 0 {
		leng = bits.Len(b)
	}
	inverse := uint(0)
	for i := 0; i < leng; i++ {
		inverse = inverse<<1 | b>>i&1
	}
	return inverse == b
}
