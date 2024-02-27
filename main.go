package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS and allow all origins
	r.Use(cors.Default())

	r.POST("/api/create", createFileOrFolder)
	r.POST("/api/list", listFiles)
	r.POST("/api/read", readFileContent)

	r.Run(":8080")
}

func createFileOrFolder(c *gin.Context) {
	var req struct {
		Path     string `json:"path"`
		IsFolder bool   `json:"isFolder"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if req.IsFolder {
		err = os.MkdirAll(req.Path, 0755)
	} else {
		_, err = os.Create(req.Path)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File or folder created successfully"})
}

func listFiles(c *gin.Context) {
	var req struct {
		Path string `json:"path"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received path for debugging
	fmt.Println("Received path:", req.Path)

	files, err := os.ReadDir(req.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{"files": fileList})
}

func readFileContent(c *gin.Context) {
	var req struct {
		Path string `json:"path"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileContent, err := os.ReadFile(req.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file content", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": string(fileContent)})
}
