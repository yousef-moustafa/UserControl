package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/yousef-moustafa/UserControl/database"
	"database/sql"
	_ "github.com/lib/pq"
)

// Represents data about a user
type user struct {
	ID int `json:",omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// Example data
var users = []user{
	{ID: 1, Name: "Yousef", Email: "yousef@gmail.com"},
}

// Declare db as a global variable to use in all functions
var db *sql.DB

func main() {
	// Initialize DB and Gin
	initDB()
	defer db.Close()
	router := gin.Default()

	// Define Routes
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUser)
	router.DELETE("/users/:id", deleteUser)
    router.Run("localhost:8080")
}

func initDB() {
	var err error
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to database successfully")
}

func getUsers(c *gin.Context) {
	// Serialize the Go Struct into pretty JSON
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	users = append(users, newUser)

	// Add new user to the database
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", newUser.Name, newUser.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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