# golang

learning go

# Table of Contents

- [Arrays](#Arrays)
- [Maps](#Maps)
- [Slices](#Slices)

## Arrays

- static
- numbered sequence of elements of the same type
- does not change in size
- used for go internals...( ? )
  > mostly used as a building block in the go

## Maps

- key value storage
- an ordered group of values of one type called the element type indexed by a set of unique keys of another type called the key type (k,v)
- powerful data structure that associates value of one type (key) with the value of another type (value)
- unordered collection of key value pairs

## Slices

descriptor of an array segment
includes pointer to the array
length of the segment
capacity(max length of the segment)

- dynamic
- built on top of an array
- holds values of the same type
- changes in size
- has a length and capacity..
- composite literals (?) are used to create a slice
- composite literal is created by having the type followed by curly braces and then putting the appropriate values in he curly brace area..
