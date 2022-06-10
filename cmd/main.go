package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Reverse(0b0001, 4)
	fmt.Printf("%02b", a)
}
