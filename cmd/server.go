package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"sneakers-app/internal/routes"
	"sneakers-app/pkg/db"
)

func init() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDatabase()
	db.SyncDatabase()
}

func main() {
	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	route.Use(cors.New(config))

	routes.InitRoutes(route)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
