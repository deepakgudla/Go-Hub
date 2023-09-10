//structs

package main

import "fmt"

type employee struct {
	name string
	designation string
	language string
}

type company struct {
	employee
	co string
	name string //embedded struct

}

func main() {
	fmt.Println("---structs---")

	employeeone := struct { //anonymous struct
        	name string
        	designation string
        	language string
} {
		name: "qwerty",
		designation: "golang-dev",
		language: "GO",
	}
	fmt.Printf("%T\n", employeeone)
	fmt.Println(employeeone)
	fmt.Printf("%v joined in a company as a %v skilled in %v\n", employeeone.name, employeeone.designation, employeeone.language)
	

	//embedded-struct
	comp := company {
		employee: employee {
			name: "qwerty",
                	designation: "golang-dev",
                	language: "GO",

		},
		name: "mnop",
		co: "Google",
	}
	fmt.Println("embedded struct:")
	fmt.Println(comp)
	fmt.Printf("%v joined at %v as a %v who is skilled in %v\n", employeeone.name, comp.co, employeeone.designation, employeeone.language)
	fmt.Printf("Hello %v\n", comp.name)
	fmt.Println(employeeone.name)
	fmt.Println(comp.employee.name) //embedded struct

}
