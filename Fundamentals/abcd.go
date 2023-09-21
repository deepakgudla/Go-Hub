/*

%s is for strings
%d for integers
%U for unicode numbers...

*/

package main

import (
	"fmt"
)

func main() {
	const name = "worldcup"
	const period = 4
	//unicode := 2693
	fmt.Printf("%s is played for every %d years.\n", name, period)
	//fmt.Println("Unicode number: %u\n", unicode);
}
