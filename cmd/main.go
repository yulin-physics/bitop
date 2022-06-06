package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.ColumnJoin([]uint{0b001101, 0b11111}, 6)

	fmt.Printf("%02b", a)
}
