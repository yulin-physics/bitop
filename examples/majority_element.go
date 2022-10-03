package examples

import (
	"github.com/yulin-physics/bitop"
)

/////////////////////////////PROBLEM/////////////////////////////////
////////////////////////////////////////////////////////////////////
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	// counter for the occurrences of each bit
	bits := make([]int, 32)
	for _, num := range nums {
		for i := range bits {
			if bitop.GetBitAtIndex(bitop.NewUnit(uint(num), 32), i) == 1 {
				bits[i]++
			}
		}
	}

	//resconstruct the number
	var number uint
	for i, count := range bits {
		if count > len(nums)/2 {
			number = bitop.FlipAtIndex(bitop.NewUnit(uint(number), 32), i)
		}
	}
	return int(number)
}

/////////////////////////////SOLUTION/////////////////////////////////
////////////////////////////////////////////////////////////////////
// If the majority element occurs more than n/2 time, then every bit that makes up the number have count greater than n/2. Thus we can reconstruct the number by creating a 32 bit number for each bit that occurs more than n/2 times.
