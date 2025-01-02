package models //storing employee data in db

type Employee struct {
	Name          string
	Designation   string
	UserID        int
	Location      string
	ContactNumber int
}
