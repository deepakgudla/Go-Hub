# golang

learning go

# Table of Contents

- [Arrays](#Arrays)
- [Maps](#Maps)
- [Slices](#Slices)
- [Pointers](#Pointers)

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

## Pointers

- type of data that refers to or points to the location in memory of a value
- go has several reference types including

  - pointers
  - slices
  - maps
  - channels
  - functions
  - INTERFACES

- holds the memory address of a value
- allows to directly access and modify the memory location of a value
  can also be called as reference types.
  pointers pass the memory address of the variable
  direct reference to the memory address
  pointer is a memory address
  every location in a pointer has an address
  all values are stored in memory
  pointers refers to a variable that holds the memory address.
  \*\*pointers allows us to directly manipulate memory and vuild complex data structures..
  improper usage of pointers can lead to bugs and errors..

* fundamental operations involving pointers..
  address operator ( & ) --> gives the address
  this operator is used to get the memory address of the var
  eg : if x := 1357 &X will give a pointer to 1357 , it is basically a memory address where the value 1357 is stored...
  dereferencing operator ( * ) --> gives the value/reference at/to memory address
  this operator is used to get the value stored at a memory address
  eg: if a is a pointer to an integer(1357) then *a gives the integer that a points to..
  if a var is a ptr to something, which is an address, a reference to some address,
  where a value is stored, we can see the value a that address using asterisk ( \* )
* POINTER itself is passed by value(fn gets a copy of the address) but the data it points to is same
  dereferencing(\* ) the pointer and modifying the value it points to inside the function will modify the original value
  In go all data is passed by value, i.e whenever we pass data to a function,
  go creates a copy of that data and assigns the copy to a parameter var
  the function can do whatever it wants to copy without affecting the original data

* when passed by value any changes made in this function wont affect the original value..

* IN GO EVERYTHING IS PASSED BY VALUE
