package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Repeat(0b01, 3, 2)
	fmt.Printf("%02b", a)
}
