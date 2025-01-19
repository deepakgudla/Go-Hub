package main

import (
	"context"
	"go-mongodb/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env var", err)
	}
	log.Println("Env var loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Database connection failed :(", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("Successfully connected to Database :) ")
}

func main() {
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	collection := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	employeeService := usecase.EmployeeService{MongoCollection: collection}

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	//routes
	r.HandleFunc("/employee", employeeService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", employeeService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", employeeService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", employeeService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", employeeService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee", employeeService.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("server is running on port 1357")
	http.ListenAndServe(":1357", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database running :) "))
}
