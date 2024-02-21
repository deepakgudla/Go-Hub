package models

type Employee struct {
  Name string
  ContactNumber int
  Designation string
  UserID int

}

var employees = []Employee{}

func getEmployees() {}

//storing employee data in db
