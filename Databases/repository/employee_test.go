package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"go-mongodb/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoClient() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URI")
	log.Println("-------------", mongoURI)
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set in the environment")
	}

	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("error while connecting to MongoDB:", err)
	}

	log.Println("MongoDB connection established")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed:", err)
	}

	log.Println("Ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := MongoClient()
	defer func() {
		if err := mongoTestClient.Disconnect(context.Background()); err != nil {
			t.Fatal("failed to disconnect MongoDB client:", err)
		}
	}()

	emp1 := uuid.New().String()

	collection := mongoTestClient.Database("company").Collection("employee_test")
	empRepo := EmployeeRepo{MongoCollection: collection}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:        "Bob",
			Age:         24,
			Designation: "Backend Developer",
			EmployeeID:  emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("failed to insert employee:", err)
		}

		t.Log("Employee created successfully:", result)
	})
}
