package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Join([]uint{0b000111100, 0b1}, 0b1)
	fmt.Printf("%02b",a)
}
