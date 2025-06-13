package repository

import (
	"context"
	"fmt"
	"go-mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

// Insert Employee
func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// Finding Employee by ID 
func (r *EmployeeRepo) FindEmployeeByID(empID string) (*model.Employee, error) {
	var emp model.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, err
}

// Finding All Employees
func (r *EmployeeRepo) FindAllEmployee() ([]model.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("decode results error %s", err.Error())
	}

	return emps, nil
}

// Updating Employee by ID
func (r *EmployeeRepo) UpdateEmployeeByID(empID string, updateEmp *model.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
		bson.D{{Key: "$set", Value: updateEmp}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

//Deleting Employee by ID
func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

// Deleting All Employees 
func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
