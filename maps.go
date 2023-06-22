//map in GO

package main

import "fmt"

func main() {
	fmt.Println("---MAPS---")
	cricket := map[string]int{
		"test": 90,
		"odi":  50,
		"t20":  20,
	}
	fmt.Println("no of overs played per day in a test match:", cricket["test"])
	fmt.Println(cricket)
	fmt.Printf("%#v\n", cricket)

	//creating maps using make
	franchiseLeagues := make(map[string]int)
	franchiseLeagues["ipl"] = 1
	franchiseLeagues["bbl"] = 2
	fmt.Println(franchiseLeagues)
	fmt.Println("length of this map", len(franchiseLeagues))

	//range of a map
	fmt.Println("---map-range---")
	for k, v := range franchiseLeagues {
		fmt.Println(k, v)
	}

	//to print only the value
	fmt.Println("printing only the value of map:")
	for _, v := range franchiseLeagues {
		fmt.Println(v)
	}

	//prinitng key of the map
	fmt.Println("printing only the key of map:")
	for k := range franchiseLeagues {
		fmt.Println(k)
	}

	//deleting an element from the map
	delete(franchiseLeagues, "bbl")

	//prints 0 if we try to print a key that does not exist
	fmt.Println("value of deleted key", franchiseLeagues["bbl"])

	//delete(franchiseLeagues, "bbl") -- wont throw an error
	//even if we try to print a key that does not exist compiler wont throw any error
	fmt.Println("values in map after deletion")
	fmt.Println(franchiseLeagues)

	//comma ok idiom helps us to find out whether or not key existed in the map

	z, ok := franchiseLeagues["ipl"]
	if ok {
		fmt.Println("value of a key is", z)
	} else {
		fmt.Println("key does not exist")
	}

	//another approach for using , ok idiom
	// if z, ok := franchiseLeagues["bbl"]; ok { //can also "!ok"
	// 	fmt.Println("value of a key is", z)

	// } else {
	// 	fmt.Println("key does not exist")
	// }

}
