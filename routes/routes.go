package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/v1/events", GetEvents)
	server.GET("/healthcheck",healthCheckHandler)
	server.GET("/v1/events/:id",GetEvent)
	server.POST("/v1/events", CreateEvent)
	server.PUT("/v1/events/:id", UpdateEvent)
	server.DELETE("/v1/events/:id",DeleteEvent)
}
