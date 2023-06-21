package main

import "fmt"

func main() {
	var sports = []string{"cricket", "badminton", "tennis"}
	fmt.Println(sports)

	sports = append(sports, "football", "volleyball", "basketball", "golf", "wrestling")
	fmt.Println(sports)
	sports = append(sports[1:5])
	fmt.Println(sports)
	sports = append(sports[2:6])
	fmt.Println(sports)
	//sports = append(sports[1:5][2:6])
	//fmt.Println(sports)
	//1:5 - 1 2 3 4
	//2:6 - 2 3 4 5
}
