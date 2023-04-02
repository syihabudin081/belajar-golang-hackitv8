package controllers

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"challenge8/db"
)



type Book struct{
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}


// CreateBook creates a new book
func CreateBook(c *gin.Context) {
    var book Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }
    defer db.Close()

    query := "INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id"
    var id int
    err = db.QueryRow(query, book.Title, book.Author).Scan(&id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }

    book.ID = id
    c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": book})
}


func GetAllBooks(c *gin.Context) {
    // Connect to the database
    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
        return
    }
    defer db.Close()

    // Execute the query
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books from database"})
        return
    }
    defer rows.Close()

    // Build the response
    var books []Book
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book details from database"})
            return
        }
        books = append(books, book)
    }
    c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
    // Get the book ID from the request parameters
    bookID := c.Param("bookID")

    // Parse the JSON request body into a Book struct
    var book Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Connect to the database
    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
        return
    }
    defer db.Close()

    // Update the book in the database
    result, err := db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3", book.Title, book.Author, bookID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
        return
    }

    // Check if the book was updated successfully
    rowsAffected, err := result.RowsAffected()
    if err != nil || rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    // Return the updated book as a JSON response
    c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func GetBookByID(c *gin.Context) {
	// Get the book ID from the request parameters
	bookID := c.Param("bookID")

	// Connect to the database
	db, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Execute the query
	var book Book
	err = db.QueryRow("SELECT * FROM books WHERE id = $1", bookID).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Return the book as a JSON response
	c.JSON(http.StatusOK, book)
}


func DeleteBook(c *gin.Context) {
	// Get the book ID from the request parameters
	bookID := c.Param("bookID")

	// Connect to the database
	db, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Delete the book from the database
	result, err := db.Exec("DELETE FROM books WHERE id = $1", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	// Check if the book was deleted successfully
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Return a 204 No Content status code
	c.JSON(http.StatusOK , gin.H{"message": "Book deleted successfully"})
}