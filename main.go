package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct user represents data of user.
type user struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DOB         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

// users slice to seed user data.
var users = []user{
	{ID: "1", Name: "Aadithya", DOB: "2002-07-13", Address: "Chennai", Description: "Root User", CreatedAt: "2022-12-29"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser user

	// Bind the received JSON to newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add newUser to the slice of users.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByID locates the user whose ID value matches
// the id parameter sent by the client, then returns
// that user as a reponse.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
