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
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUser)
	router.DELETE("/users/:id", deleteUser)
    router.Run("localhost:8080")

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

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over users slice looking for a user with id parameter value given in context
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// New slice to exclude deleted user
	var newUsers []user
	userFound := false

	for _, a := range users {
		if a.ID != id {
			newUsers = append(newUsers, a)
		} else {
			userFound = true
		}
	}
	users = newUsers
	if userFound {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

}