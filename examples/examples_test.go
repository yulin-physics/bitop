package examples

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "none",
			nums:     []int{},
			expected: -1,
		},
		{
			name:     "zero",
			nums:     []int{0, 0},
			expected: 0,
		},
		{
			name:     "one",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "some",
			nums:     []int{1, 2, 3, 3, 3},
			expected: 3,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := majorityElement(tc.nums)
			if result != tc.expected {
				t.Fatalf("[TestMajorityElement][%s]: Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}
