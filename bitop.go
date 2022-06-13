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

// GetBitAtIndex returns the bit at index `ind` of the given binary, index counting from left to right from zero as usual
// `Leng` is an optional argument to specify the exact length of the binary including leading zeroes
func GetBitAtIndex(b uint, ind int, leng int) uint {
	if leng == -1 {
		leng = bits.Len(b)
	}
	return b >> uint(leng-ind-1) & 1
}

// SplitAt returns the binary argument in two halves at the index specified [0, index), excluding leading zero bits if leng is -1
func SplitAt(b uint, index int, leng int) []uint {
	if leng < 0 {
		leng = bits.Len(b)
	}
	firstHalf := uint(0) | b>>(leng-index)
	return []uint{firstHalf, b ^ (firstHalf << (leng - index))}
}

// TruncateFromRight returns the binary truncated up to the index from the right, exclusive of the index
func TruncateFromRight(b uint, pos int) uint {
	if pos < 0 {
		return b
	}
	return b >> pos
}

// ClearFromRight returns the binary bits set to zero up to the index from the right, exclusive of the index
func ClearFromRight(b uint, index int) uint {
	if index < 0 {
		return b
	}
	return b &^ (1<<index - 1)
}

// TruncateFromLeft returns binary truncated up to the index from the left, exclusive of the index; excluding leading zero bits if leng is -1
func TruncateFromLeft(b Unit, index int) uint {
	if index < 0 {
		return b.value
	}
	return b.value &^ ((1<<index - 1) << (b.leng - index))
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
