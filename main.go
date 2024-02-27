package main

import (
	"ccs_3/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS and allow all origins
	r.Use(cors.Default())

	r.POST("/api/create", controllers.CreateFileOrFolder)
	r.POST("/api/list", controllers.ListFilesOrFolder)
	r.POST("/api/read", controllers.ReadFileContent)

	r.Run(":8080")
}
