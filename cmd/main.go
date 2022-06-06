package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.GetBitAtIndex(0b0101, 1, 4)
	fmt.Printf("%02b", a)
}
