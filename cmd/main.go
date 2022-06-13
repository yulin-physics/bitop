package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Contains(bitop.NewUnit(0b0101, 4), bitop.NewUnit(0b11, 2))
	fmt.Printf("%v", a)
}
