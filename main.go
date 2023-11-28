package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router
    r := gin.Default()

    // Define a route for the root endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to the Gin REST API!",
        })
    })

    // Define a route for an example API endpoint
    r.GET("/api/example", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "data": "This is an example API endpoint.",
        })
    })

    // Run the server on port 8080
    r.Run(":8080")
}