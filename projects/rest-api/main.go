package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("api")
	server := gin.Default() //gin - pre configured server

	server.GET("/employees", getEmployees)

	server.Run(":1357")
}

func getEmployees(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"text": "Hello World!"})
}
