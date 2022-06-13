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
