// CRUD operations using structs and slices...
package main

import (
	"encoding/json"
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
	//todo..
}

func deleteMovie(b http.ResponseWriter, a *http.Request) {
	b.Header().Set("Content-Type", "application/json")
	params := mux.Vars(a)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(b).Encode(employees)
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
	a.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 1357\n")
	log.Fatal(http.ListenAndServe(":1357", a))

}
