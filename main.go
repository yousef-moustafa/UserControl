package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/yousef-moustafa/UserControl/database"
)

// Represents data about a user
type user struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// Example data
var users = []user{
	{ID: "1", Name: "Yousef", Email: "yousef@gmail.com"},
}

func main() {
	// router := gin.Default()
    // router.GET("/albums", getUsers)
    // router.Run("localhost:8080")

	db, err := database.ConnectDB()
	if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
	fmt.Println("Connected to database successfully")
	defer db.Close()
}

func getUsers(c *gin.Context) {
	// Serialize the Go Struct into pretty JSON
	c.IndentedJSON(http.StatusOK, users)
}
