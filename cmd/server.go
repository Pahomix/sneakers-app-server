package main

import (
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

	routes.InitRoutes(route)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
