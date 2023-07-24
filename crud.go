// CRUD operations using structs and slices...
package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID string `json:"id"`
	FullName string `json:"fullname"`
	Designation string `json:"designation"`
	Company *Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
	Location string `json:"location"` 

}

var employees []Employee

func getEmployees(b http.ResponseWriter, a *http.Request) {}
func getEmployee(b http.ResponseWriter, a *http.Request) {}
func createEmployee(b http.ResponseWriter, a *http.Request) {}
func updateEmployee(b http.ResponseWriter, a *http.Request) {}
func deleteMovie(b http.ResponseWriter, a *http.Request) {}

func main() {
	fmt.Println(("crud operations....."))

	a := mux.NewRouter()

	employees = append(employees, Employee{ID: "001", FullName:"qwertyuiop", Designation: "go-dev", Company: &Company{Name: "qwerty", Location: "mnop"}})
	employees = append(employees, Employee{ID: "013", FullName:"abcdef", Designation: "rust-dev", Company: &Company{Name: "xyz", Location: "pqrs"}})
	employees = append(employees, Employee{ID: "077", FullName:"xyx", Designation: "js-dev", Company: &Company{Name: "mnop", Location: "xyz"}})



	//routes
	a.HandleFunc("/employees", getEmployees).Methods("GET")
	a.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	a.HandleFunc("/employees", createEmployee).Methods("POST")
	a.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	a.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 1357\n")
	log.Fatal(http.ListenAndServe(":1357", a))


}
