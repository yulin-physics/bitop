package bitop

import (
	"math/bits"
)

// Unit consists of a binary value which the bitop functions act on, and the number of bits in the binary as leng
type Unit struct {
	value uint
	leng  int
}

// NewUnit returns a new unit type required by most functions
// Give -ve leng to take default length measure of the input binary, which omits leading zeros
func NewUnit(b uint, leng int) Unit {
	if leng < 0 {
		leng = bits.Len(b)
	}
	return Unit{value: b, leng: leng}
}

// Contains returns true if the binary `b` has at least one section that matches with binary `sub`
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

// SplitAt returns the binary argument in two halves at the index specified [0, ind)
func SplitAt(b Unit, ind int) []uint {
	if ind < 0 {
		return []uint{b.value}
	}
	firstHalf := uint(0) | b.value>>(b.leng-ind)
	return []uint{firstHalf, b.value ^ (firstHalf << (b.leng - ind))}
}

// TruncateFromRight returns the binary truncated up to the index from the right, exclusive of the index `ind`
func TruncateFromRight(b uint, pos int) uint {
	if pos < 0 {
		return b
	}
	return b >> pos
}

// ClearFromRight returns the binary bits set to zero up to the index from the right, exclusive of the index `ind`
func ClearFromRight(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value &^ (1<<ind - 1)
}

// TruncateFromLeft returns binary truncated up to the index from the left, exclusive of the index `ind`
func TruncateFromLeft(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value &^ ((1<<ind - 1) << (b.leng - ind))
}

// RemoveBit returns the binary with the bit at index removed, length of the binary decreases by one
func RemoveBit(b Unit, ind int) uint {
	new := uint(0)
	for i := 0; i < b.leng; i++ {
		if i == ind {
			continue
		}
		new = new<<1 | b.value>>(b.leng-i-1)&1
	}
	return new
}

// Join returns a single binary by combining all binary values together separated by the given separator
func Join(bs []Unit, sep Unit) uint {
	joined := uint(0)
	for i, b := range bs {
		joined = joined<<uint(b.leng) | b.value
		if i == len(bs)-1 {
			break
		}
		joined = joined<<sep.leng | sep.value
	}
	return joined
}

// ColumnJoin joins the binary values in the array at each corresponding bit position to form columns
func ColumnJoin(rows []uint, colLeng int) []uint {
	cols := make([]uint, colLeng)
	for i := 1; i <= colLeng; i++ {
		for j := 0; j < len(rows); j++ {
			cols[i-1] = cols[i-1]<<1 | rows[j]>>(colLeng-i)&1
		}
	}
	return cols
}

// Repeat returns a binary that is a repetition of the given bit pattern for `count` number of repetitions
func Repeat(b Unit, count int) uint {
	combined := uint(0)
	for i := 0; i < count; i++ {
		combined = combined<<b.leng | b.value
	}
	return combined
}

// Replace returns a binary with any old bit pattern replaced by new, up to n times of occurrences
func Replace(b Unit, old Unit, new Unit, n int) uint {
	if n < 0 {
		return b.value
	}

	result := uint(0)
	for i := 0; i < b.leng; {
		window := TruncateFromLeft(b, i)
		window = TruncateFromRight(window, b.leng-i-old.leng)
		if window == old.value && n > 0 {
			result = result<<new.leng | new.value
			n--
			i += old.leng
		} else {
			result = result<<1 | GetBitAtIndex(b, i)
			i++
		}
	}
	return result
}

// FlipAtIndex flips the bit at the specified index in the binary
func FlipAtIndex(b Unit, ind int) uint {
	if ind < 0 {
		return b.value
	}
	return b.value ^ 1<<uint(b.leng-ind-1)
}

// Flip returns a binary with all bits flipped
func Flip(b Unit) uint {
	result := uint(0)
	for i := b.leng - 1; i >= 0; i-- {
		result = result<<1 | (1 ^ b.value>>i&1)
	}
	return result
}

// Reverse returns a binary with bits in reversed order
func Reverse(b Unit) uint {
	reversed := uint(0)
	for i := 0; i < b.leng; i++ {
		reversed = reversed<<1 | b.value>>i&1
	}
	return reversed
}

// IsPalindrome checks if the binary is symmetrical
func IsPalindrome(b Unit) bool {
	inverse := uint(0)
	for i := 0; i < b.leng; i++ {
		inverse = inverse<<1 | b.value>>i&1
	}
	return inverse == b.value
}
