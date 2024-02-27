package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateFileOrFolder(c *gin.Context) {
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
