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
