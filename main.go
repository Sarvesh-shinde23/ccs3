package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS and allow all origins
	r.Use(cors.Default())

	r.POST("/api/create", createFileOrFolder)

	r.Run(":8080")
}

func createFileOrFolder(c *gin.Context) {
	var req struct {
		RelativePath string `json:"path"`
		IsFolder     bool   `json:"isFolder"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	absolutePath := filepath.Join(".", req.RelativePath)

	var err error
	if req.IsFolder {
		err = os.MkdirAll(absolutePath, 0755)
	} else {
		_, err = os.Create(absolutePath)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File or folder created successfully"})
}
