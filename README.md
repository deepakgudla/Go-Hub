# golang

learning go

# Table of Contents

- [Arrays](#arrays)
- [Maps](#maps)
- [Slices](#slices)
- [Pointers](#pointers)
- [functions](#functions)
- [channels](#channels)

## Arrays

[arrays.go](https://github.com/deepakgudla/golang/blob/main/arrays.go)

- static
- numbered sequence of elements of the same type
- does not change in size
- used for go internals...( ? )
  - mostly used as a building block in the go
- syntax : 1. array of size x
  ```golang
  var x[5] int
  ```
- syntax -2 : array literal
  ```golang
   z := [...]string{"abcd", "mnop"}
  ```
- syntax -3: declaring an array
  ```golang
   y := [5]int{1, 2, 3, 4, 5}
  ```

## Maps

[maps.go](https://github.com/deepakgudla/golang/blob/main/maps.go)

- key value storage
- an ordered group of values of one type called the element type indexed by a set of unique keys of another type called the key type (k,v)
- powerful data structure that associates value of one type (key) with the value of another type (value)
- unordered collection of key value pairs
- syntax: here string is the key & int is the value
- when iterating over maps, the order is not deterministic
  ```golang
  <map name> := make(map[string]int) var m map[keyType]valueType
  func delete(m map[Type]Type1, key Type)  deletes the element with the key from the map.
  ``` 

## Slices

[slices.go](https://github.com/deepakgudla/golang/blob/main/slices.go)

descriptor of an array segment
includes pointer to the array
length of the segment
capacity(max length of the segment)

- dynamic
- built on top of an array
- holds values of the same type
- changes in size
- has a length and capacity.. (len(slice)), (cap(slice)) 
- composite literals (?) are used to create a slice
- composite literal is created by having the type followed by curly braces and then putting the appropriate values in he curly brace area..
- syntax :
  ```golang
  sliceName := []int{0, 1, 3, 5, 7,}
  ```
- syntax
  ```golang
  multiSlice := [][]string{weeks, days}
  ```

## Pointers

[pointers.go](https://github.com/deepakgudla/golang/blob/main/pointers.go)

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
  - address operator ( & ) - gives the address
    this operator is used to get the memory address of the var
    - eg : if x := 1357 &X will give a pointer to 1357 , it is basically a memory address where the value 1357 is stored...
  - dereferencing operator ( \* ) --> gives the value/reference at/to memory address
    this operator is used to get the value stored at a memory address
    - eg: if a is a pointer to an integer(1357) then \*a gives the integer that a points to..
      if a var is a ptr to something, which is an address, a reference to some address,
      where a value is stored, we can see the value a that address using asterisk ( \* )
* POINTER itself is passed by value(fn gets a copy of the address) but the data it points to is same
  dereferencing(\* ) the pointer and modifying the value it points to inside the function will modify the original value
  In go all data is passed by value, i.e whenever we pass data to a function,
  go creates a copy of that data and assigns the copy to a parameter var
  the function can do whatever it wants to copy without affecting the original data
* when passed by value any changes made in this function wont affect the original value..

* IN GO EVERYTHING IS PASSED BY VALUE
* syntax:

```golang
func main() {
  a := 1357
  b := &a
fmt.Println(*b) //prints the value of a
  *b = 777
fmt.Println(b) //prints the address of b
}
```

## Functions

[functions.go](https://github.com/deepakgudla/golang/blob/main/functions.go)

- can be assigned to variables
- can be passed as arguments to other functions
- can be returned as values from other functions
- when a function is assigned to a variable, the variable stores a reference to the function
  syntax -

  ```golang
  func (receiver) identifier (parameters) (returns) {}
  ```

  - func - keyword
  - receiver - type of method
  - identifier - name of the function

Sprint - prints the string

- variadic parameter

  - a function which takes unlimited number of arguments is a variadic parameter

- parameters & arguments
  - define func with parameters (if any)
  - call our func and pass in arguments
- types

  - no params, no returns
    ```golang
    func one() {
    fmt.Println("type-1: no params and no returns")
    ```
  - 1 param no returns
    ```golang
      func Two(a string) {
      fmt.Println("type-2:", a) //decalre a in main func()
    }
    ```
  - 1 param, 1 return
    ```golang
    func Three(b string) string {
    return fmt.Sprint("type-3:", b) //declare b in main func()
    }
    ```
  - 2 params, 2 returns

    ```golang
    func nameFour(format string, overs int) (string, int) {
    overs = 90
    return fmt.Sprint("no of overs played in a day in ", format), overs
    }
    ```

    ## Channels
    
    - permit communication between goroutines can be buffered or unbuffered.
