package routes

import (
	"github.com/gin-gonic/gin"
	"sneakers-app/internal/handlers/auth"
)

func InitRoutes(routes *gin.Engine) {
	routes.POST("/login", auth.LoginHandler)
	routes.POST("/register", auth.RegisterHandler)

}
