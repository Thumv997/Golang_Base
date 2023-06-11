package routes

import (
	// "log"

	"lore_project/internal/handles"
	"lore_project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, service *services.UserService) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", handles.NewAuthHandler(service).Login)
		authRoutes.POST("/register", handles.NewAuthHandler(service).Register)
		// authRoutes.POST("/logout", logoutHandler)
	}
}
