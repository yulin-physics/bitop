package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.RemoveBit(0b1111000, 3)
	fmt.Printf("%2b", a)
}
