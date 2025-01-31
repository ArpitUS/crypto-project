package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple hashing function (uses Rust)
func hashHandler(c *gin.Context) {
	data := c.Query("data")
	if data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data parameter"})
		return
	}

	// Call Rust FFI function (later integration)
	hashed := fmt.Sprintf("Hash of %s", data)
	c.JSON(http.StatusOK, gin.H{"hashed": hashed})
}

func main() {
	r := gin.Default()
	r.GET("/hash", hashHandler)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
