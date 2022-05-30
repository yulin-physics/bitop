package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.TruncateFromLeft(0b111100, 3)
	fmt.Printf("%2b", a)
}
