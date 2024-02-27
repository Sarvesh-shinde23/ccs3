package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ReadFileContent(c *gin.Context) {
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
