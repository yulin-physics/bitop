package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Contains(0b11110, 0b00, 6, 2)
	fmt.Printf("%v", a)
}
