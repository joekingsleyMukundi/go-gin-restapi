package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{Id: "1", Title: "sailing sun set", Author: "Khali", Quantity: 3},
	{Id: "2", Title: "Docking sun set", Author: "Joe", Quantity: 6},
	{Id: "3", Title: "Seting sun set", Author: "Mukundi", Quantity: 15},
}

func getBookById(id string) (*Book, error) {
	for i, b := range Books {
		if b.Id == id {
			return &Books[i], nil
		}
	}
	return nil, errors.New("no book was found with that id")
}
func deleteBookByIdFunc(id string) (err error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	Books = append(Books[:i], Books[i+1:]...)
	return nil
}
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}
func getBook(c *gin.Context) {
	// to geT the dynamic parameter we do this
	id := c.Param("id")
	//NOTE  for query parameters we do this
	//  id, err := c.GetQuery("id")

	book, err := getBookById(id)
	if err != nil {
		// gin.H helps us write json or rather helsp us just create ajson format of this type quickly

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Books = append(Books, newBook)
	book, err := getBookById(newBook.Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusCreated, book)
}
func deleteBookById(c *gin.Context) {
	id := c.Param("id")
	err := deleteBookByIdFunc(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
func main() {

	// set up router with gin just like route in node js
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("book/:id", getBook)
	router.DELETE("book/:id", deleteBookById)
	// for patch request we just sat router.PATCH

	router.Run("localhost:8080")

}
