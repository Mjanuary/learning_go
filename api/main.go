package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
	ID 			string `json:"id"`
	Title 		string `json:"title"`
	Author 		string `json:"author"`
	Quantity 	int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In search of lost time", Author: "Marcel Pro", Quantity: 3},
	{ID: "2", Title: "New stuff", Author: "James Pro", Quantity: 4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router("localhost:8080")
}