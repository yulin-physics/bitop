package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Flip(0b0101, 4)
	fmt.Printf("%02b", a)
}
