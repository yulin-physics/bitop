# Bit Operations in Go

This bit operations library mimicks the functions of Go strings library.

For learning purposes only, use at your own risk.

## Index

Types:

[Unit](#types)

Functions:

[ClearFromRight](#func-clearfromright)
[Contains](#func-contains)
[ColumnJoin](#func-columnjoin)
[GetBitAtIndex](#func-getbitatindex)
[Join](#func-join)
[LastIndex](#func-lastindex)
[RemoveBit](#func-removebit)
[Repeat](#func-repeat)
[SplitAt](#func-splitat)
[TruncateFromLeft](#func-truncatefromleft)
[TruncateFromRight](#func-truncatefromright)

## Types

The visible length of the binary is required in bit shifting in this package, therefore a new type `Unit` is used.

An Unit composed of the binary itself and length of the binary.

```
type Unit struct {
	value uint
	leng  int
}
```

## Functions

### func Contains

`func Contains(b, sub Unit) bool`

Checks the target binary `b` has at least one part that matches the sub-binary value.

`func IsPalindrome(b uint, leng int) bool`

Returns true if the binary contains symmetry.

### func LastIndex

`func LastIndex(b, sub Unit) int`

Finds the index (counting from left to right) of the last bit pattern in `b` that matches `sub`.

### func GetBitAtIndex

`func GetBitAtIndex(b Unit, ind int) uint`

Finds the bit at index `ind` of binary `b`.

### func SplitAt

`func SplitAt(b Unit, index int) []uint`

Splits the binary in two at the index, returns the 2 sub-binaries.

### func Join

`func Join(bs []Unit, sep Unit) uint`

Join combines the binary values into one, separated by the given delimiter.

### func ColumnJoin

`func ColumnJoin(rows []uint, colLeng int)`

ColumnJoin combines the binary values in each corresponding bit position, forming array of columns.

colLeng is usually the bit length of an element in rows, but since leading zeroes are ommited and in case of variable length binaries in input, user needs to specify the bit length.

### func TruncateFromRight

`func TruncateFromRight(b uint, pos int) uint`

TruncateFromRight trims off bits from right, up to but not including the index.

### func TruncateFromLeft

`func TruncateFromLeft(b Unit, ind int) uint`

TruncateFromLeft trims off bits from left, up to but not including the index.

### func ClearFromRight

`func ClearFromRight(b Unit, ind int) uint`

Clear preserves the binary length and resets the bits from right to zero, up to but not including the index.

### func RemoveBit

`func RemoveBit(b Unit, index int) uint`

RemoveBit removes the bit at the index from the binary.

### func Repeat

`func Repeat(b uint, leng int, count int)`

Repeat constructs a binary based on a given repeating bit pattern.

### Replace

`func Replace(b uint, old uint, new uint, n int, leng int, oldLeng int, newLeng int)`

Replace constructs a new binary with the old sub bits replaced by the new up to n times

## Flip

`func FlipAtIndex(b uint, index int, leng int)`

Flips the bit at index i in the binary

`func Flip(b uint, leng int)`

Flips all bits

## Reverse

`func Reverse(b uint, leng int)`

Returns bits in reversed order

## Resources

- [How to Use a Private Go Module in Your Own Project](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project)
- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules)
