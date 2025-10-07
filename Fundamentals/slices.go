package main

import "fmt"

type Slice_ struct{}

func (s Slice_) Name() string {
	return "Slice"
}

func (s Slice_) Run() {

	//unfurling a slice
	fmt.Println("unfurling a slice")
	c := []int{2, 4, 6, 8, 10, 12}
	b := add(c...)
	fmt.Println("the sum is ", b)

	sports := []string{"cricket", "badminton", "tennis"}
	fmt.Println(sports)

	//to append values to a slice,"append function is used"
	//append func returns values of same type
	//adding elements to the existing slice
	sports = append(sports, "football", "volleyball", "basketball", "golf", "wrestling")
	fmt.Println(sports)

	sports = sports[1:5] //prints 1 2 3 4
	fmt.Println(sports)

	sports = (sports[2:6]) // prints 2 3 4 5
	fmt.Println(sports)
	//sports = append(sports[1:5][2:6]) // ?
	//fmt.Println(sports)
	fmt.Println("_________________")

	a := []string{"hello", "gopher"}
	fmt.Println(a)

	for i, v := range sports {
		fmt.Printf("index %v is storing the value %v\n", i, v)
	}
	//to show just a blank value with no index
	for _, v := range sports { //purpose of underscore is to let to know the compiler that there are two values and one is blank...
		fmt.Printf("%v\n", v)
	}

	fmt.Println("accessing values by index position:")

	//accessing values by the index position
	fmt.Println(sports[0])
	fmt.Println(sports[1])
	//fmt.Println(sports[7])
	fmt.Println("length of the slice: ", len(sports))

	//printing all the values in slice using for loop
	fmt.Println("printing using for loop:")
	for i := 0; i < len(sports); i++ {
		fmt.Println(sports[i])
	}

	//slicing a slice
	//we can cut parts of the slice away with the ":" operator
	fmt.Println("slicing-a-slice")

	/* [inclusive:exclusive]
	in this case the exclusive element gets sliced...
	*/

	slicing := []int{0, 1, 3, 5, 7, 11, 13, 135, 1357, 77}
	fmt.Printf("%v\n", slicing)
	//fmt.Printf("%#v\n",slicing) --> prints the datatype..

	//[inclusive:exclusive] slicing
	fmt.Printf("ie sliced elements: %v\n", slicing[0:3])

	//[inclusive:] slice - prints including from that element
	fmt.Printf("inclusive slice: %v\n", slicing[3:])

	//[:exclusive] slice - prints the values before that position
	fmt.Printf("exclusive slice: %v\n", slicing[:5])

	//using : operator
	fmt.Printf("slicing using colon: %v\n", slicing[:])

	//deleting a slice
	slicing = append(slicing[:2], slicing[3:]...)
	fmt.Printf("deleted elements: %v\n", slicing)
	slicing = append(slicing[:1], slicing[7:]...)
	fmt.Printf("multiple deletion: %v\n", slicing) //elements from index 1-7 will be deleted

	//make
	tokens := []string{"20", "721", "1155"}
	fmt.Printf("elements: %v length: %v\n", tokens, len(tokens))
	//above is the standard implementation of slice
	//but with make we can declare length & capacity and later we can append the values in that slice

	days := make([]string, 0, 6)
	fmt.Println(days)
	fmt.Println(len(days))
	fmt.Println(cap(days))
	days = append(days, "mon", "tue", "wed", "thurs", "fri", "sat", "sun", "weekend", "abcd", "pqrs", "mnop")
	fmt.Printf("appended elements: %v\n", days)
	fmt.Println(len(days))
	fmt.Println(cap(days))

	//multi dimensional slices
	multi := [][]string{tokens, days}
	fmt.Println(multi)

	m := []int{2, 4, 6, 8}
	//n := m
	//using copy function
	n := make([]int, 4)
	copy(n, m) //(destination, source)
	fmt.Println("slice-m:", m)
	fmt.Println("slice-n:", n)
	fmt.Println("changing position index 0")
	m[0] = 100
	fmt.Println("slice-m:", m)
	fmt.Println("slice-n:", n)

	//for range in a slice
	fmt.Println("range of a slice with index and value:")
	for i, v := range days {
		fmt.Println(i, v)
	}
	//to print only value --> for _, v and print v
	//to print only index --> for i and print i
}

// unfurling a slice
func add(a ...int) int {
	fmt.Println(a)

	i := 0
	for _, v := range a {
		i += v
	}
	return i
}

func init() {
	Register(Slice_{})
}
