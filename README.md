# Bit Operations in Go

This bit operations library mimicks the functions of Go strings library.

For learning purposes only, use at your own risk.

## Types

The visible length of the binary is required in bit shifting in this package, therefore a new type `Unit` is used.

An Unit composed of the binary itself and length of the binary.

```
type Unit struct {
	value uint
	leng  int
}
```

## Index

[Contains](#func-contains)
[LastIndex](#func-lastindex)
[GetBitAtIndex](#func-getbitatindex)
[SplitAt](#func-splitat)

## Functions

### func Contains

`func Contains(b, sub Unit) bool`

Checks the target binary `b` has at least one part that matches the sub-binary value.

`func IsPalindrome(b uint, leng int) bool`

Returns true if the binary contains symmetry

### func LastIndex

`func LastIndex(b, sub Unit) int`

Finds the index (counting from left to right) of the last bit pattern in `b` that matches `sub`.

### func GetBitAtIndex

`func GetBitAtIndex(b Unit, ind int) uint`

Finds the bit at index `ind` of binary `b`.

## func SplitAt

`func SplitAt(b Unit, index int) []uint`

Splits the binary in two at the index, returns the 2 sub-binaries

## Join

`func Join(bs []uint, sep uint)`

Join combines the binary values into one, separated by the given delimiter

`func ColumnJoin(rows []uint, colLen int)`

ColumnJoin combines the binary values in each corresponding bit position, forming array of columns

colLen is usually the bit length of an element in rows, but since leading zeroes are ommited, user needs to specify the bit length

## Trim

`func TruncateFromRight(b uint, index int) `

`func TruncateFromLeft(b uint, index int)`

Truncate functions trim off bits from left (or right), up to but not including the index

`func ClearFromRight(b uint, index int) `

Clear preserves the binary length and resets the bits from right to zero, up to but not including the index

`func RemoveBit(b uint, index int)`

RemoveBit removes the bit at the index from the binary

## Repeat

`func Repeat(b uint, leng int, count int)`

Repeat constructs a binary based on a given repeating bit pattern

## Replace

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
