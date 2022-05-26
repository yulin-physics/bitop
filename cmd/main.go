package main

import (
	"fmt"

	"github.com/yulin-physics/bitop"
)

func main() {
	a := bitop.LastIndexOfZero(0b000111100)
	fmt.Println(a)
}
