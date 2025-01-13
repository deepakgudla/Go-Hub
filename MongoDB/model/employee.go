package model

type Employee struct {
	EmployeeID  string `json:"employee_id,omitempty" bson:"employee_id"`
	Name        string `json:"name,omitempty" bson:"name"`
	Age         int    `json:"age,omitempty" bson:"age"`
	Designation string `json:"designation,omitempty" bson:"designation"`
}
