package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaustubsingh09/go-auth-gin/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)

	authenticated.POST("/events", postEvents)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/event/:id", getUniqueEvent)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}
