package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.FlipAtIndex(0b101010, 2, 6)
	fmt.Printf("%02b", a)
}
