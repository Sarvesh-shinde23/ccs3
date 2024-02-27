package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ListFilesOrFolder(c *gin.Context) {
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
