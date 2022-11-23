package controllers

import (
	"net/http"

	"instrumented-app/models"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/trace"
)

var meter = global.Meter("bookstore-meter")

var queryBooksCount, _ = meter.SyncInt64().Counter(
	"bookstore/query_all_books_request_counts",
	instrument.WithDescription("The number of query all books requests received"),
)
var queryBookCount, _ = meter.SyncInt64().Counter(
	"bookstore/query_one_book_request_counts",
	instrument.WithDescription("The number of query one book requests received"),
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books
// Find all books
func FindBooks(c *gin.Context) {
	var books []models.Book

	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("controller", "books"))
	span.AddEvent("This is a sample event", trace.WithAttributes(attribute.Int("pid", 4328), attribute.String("sampleAttribute", "Test")))

	models.DB.WithContext(ctx).Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})

	queryBooksCount.Add(ctx, 1)
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.WithContext(c.Request.Context()).Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})

	queryBookCount.Add(c.Request.Context(), 1)
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.WithContext(c.Request.Context()).Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.WithContext(c.Request.Context()).Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.WithContext(c.Request.Context()).Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.WithContext(c.Request.Context()).Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
