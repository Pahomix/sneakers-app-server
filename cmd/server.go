package cmd

import (
	"github.com/gin-gonic/gin"
	"sneakers-app/internal/routes"
	"sneakers-app/pkg/db"
)

func init() {
	db.ConnectDatabase()
	db.SyncDatabase()
}

func main() {
	route := gin.Default()

	routes.InitRoutes(route)
	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
