package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.Replace(0b101010, 0b01, 0b1, 1, 6, 1, 1)
	fmt.Printf("%02b", a)
}
