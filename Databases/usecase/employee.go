package usecase

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Date  interface{}
	Error string
}

// todo
func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request)     {}
func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request)    {}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request)     {}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request)  {}
