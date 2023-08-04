package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Name: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Name: "Cloud Native Go", Author: "M.-L. Reimer, K. Köhler", Quantity: 5},
	{ID: "3", Name: "To kill a mockingbird", Author: "Harper Lee", Quantity: 3},
	{ID: "4", Name: "The Lord of the Rings", Author: "J. R. R. Tolkien", Quantity: 1},
	{ID: "5", Name: "The Da Vinci Code", Author: "Dan Brown", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func main() {
	router := gin.Default()

}
