package bitop

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		sub      Unit
		expected bool
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			sub:      NewUnit(0b11, -1),
			expected: true,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			sub:      NewUnit(0b11, -1),
			expected: false,
		},
		{
			name:     "0b010101",
			b:        NewUnit(0b010101, 6),
			sub:      NewUnit(0b01, 2),
			expected: true,
		},
		{
			name:     "0b110110",
			b:        NewUnit(0b110110, 6),
			sub:      NewUnit(0b111, 3),
			expected: false,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Contains(tc.b, tc.sub)
			if result != tc.expected {
				t.Fatalf("[TestContains][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		sub      Unit
		expected int
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			sub:      NewUnit(0b11, -1),
			expected: 4,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			sub:      NewUnit(0b11, -1),
			expected: -1,
		},
		{
			name:     "0b010101",
			b:        NewUnit(0b010101, 6),
			sub:      NewUnit(0b01, 2),
			expected: 4,
		},
		{
			name:     "0b110110",
			b:        NewUnit(0b110110, 6),
			sub:      NewUnit(0b111, 3),
			expected: -1,
		},
		{
			name:     "bit one",
			b:        NewUnit(0b110110, 6),
			sub:      NewUnit(0b1, -1),
			expected: 4,
		},
		{
			name:     "bit zero",
			b:        NewUnit(0b110110, 6),
			sub:      NewUnit(0b0, 1),
			expected: 5,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := LastIndex(tc.b, tc.sub)
			if result != tc.expected {
				t.Fatalf("[TestLastIndex][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestGetBitAtIndex(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    1,
			expected: 0b1,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			index:    4,
			expected: 0b0,
		},
		{
			name:     "0b010101",
			b:        NewUnit(0b010101, 6),
			index:    0,
			expected: 0b0,
		},
		{
			name:     "0b110110",
			b:        NewUnit(0b110110, 6),
			index:    0,
			expected: 0b1,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := GetBitAtIndex(tc.b, tc.index)
			if result != tc.expected {
				t.Fatalf("[TestGetBitAtIndex][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestSplitAt(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected []uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    1,
			expected: []uint{0b1, 0b11111},
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			index:    4,
			expected: []uint{0b0000, 0b00},
		},
		{
			name:     "0b010101",
			b:        NewUnit(0b010101, 6),
			index:    0,
			expected: []uint{0b0, 0b010101},
		},
		{
			name:     "0b110110",
			b:        NewUnit(0b110110, 6),
			index:    5,
			expected: []uint{0b11011, 0b0},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := SplitAt(tc.b, tc.index)
			for i, r := range result {
				if r != tc.expected[i] {
					t.Fatalf("[TestSplitAt][%s]: Got %v, expected %v", tc.name, result, tc.expected)
				}
			}
		})
	}
}

func TestTruncateFromRight(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        uint
		pos      int
		expected uint
	}{
		{
			name:     "ones",
			b:        0b111111,
			pos:      2,
			expected: 0b1111,
		},
		{
			name:     "zeroes",
			b:        0b0000,
			pos:      1,
			expected: 0b000,
		},
		{
			name:     "010101",
			b:        0b010101,
			pos:      3,
			expected: 0b010,
		},
		{
			name:     "1001",
			b:        0b1001,
			pos:      0,
			expected: 0b1001,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := TruncateFromRight(tc.b, tc.pos)
			if result != tc.expected {
				t.Fatalf("[TestTruncateFromRight][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestTruncateFromLeft(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    2,
			expected: 0b1111,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b0000, 4),
			index:    1,
			expected: 0b000,
		},
		{
			name:     "010101",
			b:        NewUnit(0b010101, 6),
			index:    3,
			expected: 0b101,
		},
		{
			name:     "1001",
			b:        NewUnit(0b1001, -1),
			index:    0,
			expected: 0b1001,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := TruncateFromLeft(tc.b, tc.index)
			if result != tc.expected {
				t.Fatalf("[TestTruncateFromLeft][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestClearFromRight(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    2,
			expected: 0b111100,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b0000, 4),
			index:    1,
			expected: 0b0000,
		},
		{
			name:     "010101",
			b:        NewUnit(0b010101, 6),
			index:    3,
			expected: 0b010000,
		},
		{
			name:     "1001",
			b:        NewUnit(0b1001, -1),
			index:    0,
			expected: 0b1001,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := ClearFromRight(tc.b, tc.index)
			if result != tc.expected {
				t.Fatalf("[TestClearFromRight][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestRemoveBit(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    2,
			expected: 0b11111,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b0000, 4),
			index:    1,
			expected: 0b000,
		},
		{
			name:     "010101",
			b:        NewUnit(0b010101, 6),
			index:    3,
			expected: 0b01001,
		},
		{
			name:     "1001",
			b:        NewUnit(0b1001, -1),
			index:    0,
			expected: 0b001,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := RemoveBit(tc.b, tc.index)
			if result != tc.expected {
				t.Fatalf("[TestRemoveBit][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		bs       []Unit
		sep      Unit
		expected uint
	}{
		{
			name:     "ones",
			bs:       []Unit{NewUnit(0b111111, -1), NewUnit(0b11, -1)},
			sep:      NewUnit(0b0, 0),
			expected: 0b11111111,
		},
		{
			name:     "zeroes",
			bs:       []Unit{NewUnit(0b00, 2), NewUnit(0b000, 3)},
			sep:      NewUnit(0b0, 0),
			expected: 0b00000,
		},
		{
			name:     "single separator 0b1",
			bs:       []Unit{NewUnit(0b00, 2), NewUnit(0b000, 3)},
			sep:      NewUnit(0b1, 0),
			expected: 0b001000,
		},
		{
			name:     "multiple separator 0b1",
			bs:       []Unit{NewUnit(0b1011, -1), NewUnit(0b101, -1), NewUnit(0b111, -1), NewUnit(0b0000, 4)},
			sep:      NewUnit(0b0, 1),
			expected: 0b10110101011100000,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Join(tc.bs, tc.sep)
			if result != tc.expected {
				t.Fatalf("[TestJoin][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestColumnJoin(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		rows     []uint
		colLeng  int
		expected []uint
	}{
		{
			name:     "ones",
			rows:     []uint{0b111, 0b111, 0b111, 0b111},
			colLeng:  3,
			expected: []uint{0b1111, 0b1111, 0b1111},
		},
		{
			name:     "zeroes",
			rows:     []uint{0b000, 0b000, 0b000, 0b000, 0b000},
			colLeng:  3,
			expected: []uint{0b00000, 0b00000, 0b00000},
		},
		{
			name:     "same lengths",
			rows:     []uint{0b1010, 0b0101, 0b1110, 0b0111, 0b1100},
			colLeng:  4,
			expected: []uint{0b10101, 0b01111, 0b10110, 0b01010},
		},
		{
			name:     "variable lengths",
			rows:     []uint{0b1010, 0b010, 0b110, 0b01011, 0b1100},
			colLeng:  4,
			expected: []uint{0b10011, 0b00101, 0b11110, 0b00010},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := ColumnJoin(tc.rows, tc.colLeng)
			for i, r := range result {
				if r != tc.expected[i] {
					t.Fatalf("[TestColumnJoin][%s]: Got %v, expected %v", tc.name, result, tc.expected)
				}
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		count    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			count:    2,
			expected: 0b111111111111,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b0000, 4),
			count:    0,
			expected: 0b0,
		},
		{
			name:     "010101",
			b:        NewUnit(0b010101, 6),
			count:    0,
			expected: 0b0,
		},
		{
			name:     "1001",
			b:        NewUnit(0b1001, -1),
			count:    1,
			expected: 0b1001,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Repeat(tc.b, tc.count)
			if result != tc.expected {
				t.Fatalf("[TestRepeat][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		old      Unit
		new      Unit
		n        int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			old:      NewUnit(0b1, -1),
			new:      NewUnit(0b0, -1),
			n:        2,
			expected: 0b1111,
		},
		{
			name:     "zeros",
			b:        NewUnit(0b000000, 6),
			old:      NewUnit(0b0, 1),
			new:      NewUnit(0b1, 1),
			n:        2,
			expected: 0b110000,
		},
		{
			name:     "no match",
			b:        NewUnit(0b000000, 6),
			old:      NewUnit(0b1, 1),
			new:      NewUnit(0b0, 1),
			n:        2,
			expected: 0b000000,
		},
		{
			name:     "complex match",
			b:        NewUnit(0b1010101, -1),
			old:      NewUnit(0b101, -1),
			new:      NewUnit(0b1, 1),
			n:        3,
			expected: 0b101,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Replace(tc.b, tc.old, tc.new, tc.n)
			if result != tc.expected {
				t.Fatalf("[TestReplace][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestFlipAtIndex(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		index    int
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			index:    2,
			expected: 0b110111,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			index:    3,
			expected: 0b000100,
		},
		{
			name:     "mixed",
			b:        NewUnit(0b01100101, 8),
			index:    0,
			expected: 0b11100101,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := FlipAtIndex(tc.b, tc.index)
			if result != tc.expected {
				t.Fatalf("[TestFlipAtIndex][%s]: Got %02b, expected %02b", tc.name, result, tc.expected)
			}
		})
	}
}

func TestFlip(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			expected: 0b000000,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			expected: 0b111111,
		},
		{
			name:     "mixed",
			b:        NewUnit(0b01100101, 8),
			expected: 0b10011010,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Flip(tc.b)
			if result != tc.expected {
				t.Fatalf("[TestFlip][%s]: Got %02b, expected %02b", tc.name, result, tc.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        Unit
		expected uint
	}{
		{
			name:     "ones",
			b:        NewUnit(0b111111, -1),
			expected: 0b111111,
		},
		{
			name:     "zeroes",
			b:        NewUnit(0b000000, 6),
			expected: 0b000000,
		},
		{
			name:     "mixed",
			b:        NewUnit(0b01100101, 8),
			expected: 0b10100110,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Reverse(tc.b)
			if result != tc.expected {
				t.Fatalf("[TestReverse][%s]: Got %02b, expected %02b", tc.name, result, tc.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		b        uint
		leng     int
		expected bool
	}{
		{
			name:     "ones",
			b:        0b111111,
			leng:     -1,
			expected: true,
		},
		{
			name:     "zeroes",
			b:        0b0000,
			leng:     4,
			expected: true,
		},
		{
			name:     "010101",
			b:        0b010101,
			leng:     6,
			expected: false,
		},
		{
			name:     "1001",
			b:        0b1001,
			leng:     -1,
			expected: true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IsPalindrome(tc.b, tc.leng)
			if result != tc.expected {
				t.Fatalf("[TestIsPalindrome][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}
