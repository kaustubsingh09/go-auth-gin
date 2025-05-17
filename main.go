package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaustubsingh09/go-auth-gin/db"
	"github.com/kaustubsingh09/go-auth-gin/routes"
)

func main() {
	server := gin.Default()
	fmt.Println("Server started!")
	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":8000")
}
