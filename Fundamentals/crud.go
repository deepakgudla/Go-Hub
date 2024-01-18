// CRUD operations using structs and slices...
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID          string   `json:"id"`
	FullName    string   `json:"fullname"`
	Designation string   `json:"designation"`
	Company     *Company `json:"company"`
}

type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

var employees []Employee

var ErrEmployeeNotFound = errors.New("employee not found")

func getEmployees(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	json.NewEncoder(b).Encode(employees)
}
func createEmployee(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(a.Body).Decode(&employee)
	employee.ID = strconv.Itoa(rand.Intn(1000000))
	employees = append(employees, employee)
	json.NewEncoder(b).Encode(employee)

}
func updateEmployee(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	params := mux.Vars(a)
	employeeFound := false
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(a.Body).Decode(&employee)
			employee.ID = params["id"]
			employees = append(employees, employee)
			json.NewEncoder(b).Encode(employee)
			employeeFound = true
			return
		}

		if !employeeFound {
			http.Error(b, "could not find employee or deleted", http.StatusNotFound)
		}
	}
}

func deleteEmployee(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	params := mux.Vars(a)
	employeeFound := false
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			json.NewEncoder(b).Encode(employees)
			employeeFound = true
			return
		}
		if !employeeFound {
			http.Error(b, "could not find employee or deleted", http.StatusNotFound)
			fmt.Println(http.Error)
		}
	}
}

func getEmployee(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	params := mux.Vars(a)
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(b).Encode(item)
			return
		}
	}
}

func main() {
	fmt.Println(("crud operations....."))

	a := mux.NewRouter()

	employees = append(employees, Employee{ID: "001", FullName: "qwertyuiop", Designation: "go-dev", Company: &Company{Name: "qwerty", Location: "mnop"}})
	employees = append(employees, Employee{ID: "013", FullName: "abcdef", Designation: "rust-dev", Company: &Company{Name: "xyz", Location: "pqrs"}})
	employees = append(employees, Employee{ID: "077", FullName: "xyx", Designation: "js-dev", Company: &Company{Name: "mnop", Location: "xyz"}})

	//routes
	a.HandleFunc("/employees", getEmployees).Methods("GET")
	a.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	a.HandleFunc("/employees", createEmployee).Methods("POST")
	a.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	a.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

	fmt.Printf("starting server at port 1357\n")
	log.Fatal(http.ListenAndServe(":1357", a))
}

//need to complete error handling
