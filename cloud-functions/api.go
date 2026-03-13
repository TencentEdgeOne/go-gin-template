package main

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// REST API v1 group
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", helloHandler)
		v1.GET("/health", healthHandler)

		// Users group
		users := v1.Group("/users")
		{
			users.GET("", listUsersHandler)
			users.GET("/:id", getUserHandler)
			users.POST("", createUserHandler)
		}

		// Posts group
		posts := v1.Group("/posts")
		{
			posts.GET("", listPostsHandler)
			posts.GET("/:id", getPostHandler)
		}
	}

	r.Run(":9000")
}

// helloHandler returns a welcome message
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Gin on EdgeOne Pages!",
	})
}

// healthHandler returns service health status
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"runtime":   "go",
		"version":   runtime.Version(),
		"framework": "gin",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Sample data
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 42, Name: "Charlie", Email: "charlie@example.com"},
}

// listUsersHandler returns all users
func listUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": len(users),
	})
}

// getUserHandler returns a user by ID (dynamic param)
func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found", "id": id})
}

// createUserHandler creates a new user (JSON binding)
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func createUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := User{
		ID:    len(users) + 1,
		Name:  req.Name,
		Email: req.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    newUser,
	})
}

// Posts data
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []Post{
	{ID: 1, Title: "Getting Started with Gin", Body: "Gin is a high-performance HTTP web framework written in Go."},
	{ID: 7, Title: "Deploy Go to EdgeOne", Body: "Learn how to deploy Go applications to EdgeOne Pages."},
	{ID: 3, Title: "Gin Middleware Guide", Body: "Middleware in Gin provides a way to process requests globally."},
}

// listPostsHandler returns all posts
func listPostsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"total": len(posts),
	})
}

// getPostHandler returns a post by ID
func getPostHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	for _, p := range posts {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "post not found", "id": id})
}
