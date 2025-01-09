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

	collection := mongoTestClient.Database("company").Collection("employee_test")

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

		t.Log("Suucessfully inserted 1 employee", result)
	})
}
