package bitop

import (
	"fmt"
)

// concatenate binary A B at each corresponding bit
// visualise each bit as a cell in row, concatenate each column
// 0b11111111
// 0b00000000
// -----------
// 0b10 0b10 0b10 0b10 0b10 0b10 0b10 0b10 expected

const (
	A       = 0b11111111
	B       = 0b00000000
	COL_NUM = 8
)

func main() {
	rows := []uint64{A, B}
	cols := make([]uint64, COL_NUM)
	for i := 1; i <= COL_NUM; i++ {
		for j := 0; j < len(rows); j++ {
			// cols[i-1]<<1 to make space for the new bit
			// (COL_NUM-i) to get individual bit from LEFT to RIGHT
			// &1 to get the new bit
			cols[i-1] = cols[i-1]<<1 | rows[j]>>(COL_NUM-i)&1
		}
	}
	fmt.Printf("%02b\n", cols)
}
