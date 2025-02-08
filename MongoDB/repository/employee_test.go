package repository

import (
	"context"
	"go-mongodb/model"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// testing Database connection
func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://admin:deepakgudla@cluster0.hhsh2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("error while connecting MongoDB", err)
	}

	log.Println("Successfully connected to MongoDB :) ")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	empOne := uuid.New().String()

	collection := mongoTestClient.Database("company").Collection("employees")

	empRepo := EmployeeRepo{MongoCollection: collection}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			EmployeeID:  empOne,
			Name:        "bob",
			Age:         24,
			Designation: "Backend Engineer",
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insertion of 1 operation failed", err)
		}

		t.Log("Successfully inserted 1 employee", result)
	})

	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(empOne)
		if err != nil {
			t.Fatal("Get operation failed", err)
		}

		t.Log("emp 1", result.Name)
	})
	
	//test: get all employees 
	t.Run("Get All employees", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("failed while fetching all employees", err)
		}

		t.Log("employees", results)
	})

	t.Run("Update Employee 1 Name", func(t *testing.T) {
		emp := model.Employee{
			EmployeeID:  empOne,
			Name:        "Bob The Builder",
			Age:         24,
			Designation: "Backend Engineer",
		}

		results, err := empRepo.UpdateEmployeeByID(empOne, &emp)
		if err != nil {
			log.Fatal("failed to update employee details", err)
		}

		t.Log("update count", results)
	})

	t.Run("Get Employee 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(empOne)
		if err != nil {
			t.Fatal("failed to fetch employee details", err)
		}

		t.Log("emp 1", result.Name)
	})

	t.Run("Delete Employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(empOne)
		if err != nil {
			t.Fatal("failed to delete employee ", err)
		}

		t.Log("delete count", result)
	})

	t.Run("Delete All Employees", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			t.Fatal("failed to delete all employees", err)
		}

		t.Log("Successfully deleted all employees", result)
	})
}
