package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.ClearFromRight(0b111111100, 3)
	fmt.Printf("%2b", a)
}
