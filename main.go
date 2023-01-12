package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `JSON:"id"`
	Title    string `JSON:"title"`
	Author   string `JSON:"author"`
	Quantity int    `JSON:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of the lost time", Author: "Marcel Proust", Quantity: 1},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgernald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 2},
	{ID: "4", Title: "The Man Searching for Peace", Author: "Victor Frankel", Quantity: 6},
	{ID: "5", Title: "The Girl in Room 105", Author: "Chetan Bhagat", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookbyID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookbyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookbyID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book Not Found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing ID query parameter"})
		return
	}
	book, err := getBookbyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing ID query parameter"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing ID query parameter"})
		return
	}
	book, err := getBookbyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing ID query parameter"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	// setting up router responsible for handling out different routes, and different endpoints
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookbyID)
	router.POST("/books", getBooks)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")

}
