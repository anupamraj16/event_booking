package routes

import (
	"github.com/anupamraj16/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/registrations", getRegistrations)
	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	// server.POST("/events", middlewares.Authenticate, createEvent)
	authenticatedRoutes.POST("/events", createEvent)
	authenticatedRoutes.PUT("/events/:id", updateEvent)
	authenticatedRoutes.DELETE("/events/:id", deleteEvent)
	authenticatedRoutes.POST("/events/:id/register", registerForEvent)
	authenticatedRoutes.DELETE("/events/:id/register", cancelRegistration)
	server.POST("/signup", createUser)
	server.POST("/login", login)
}
