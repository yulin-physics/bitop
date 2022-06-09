package bitop

import (
	"testing"
)

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
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IsPalindrome(tc.b, tc.leng)
			if result != tc.expected {
				t.Fatalf("[TestIsPalindrome][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}
