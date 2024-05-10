package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sneakers-app/internal/routes"
	"sneakers-app/pkg/db"
	"time"
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
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AddAllowHeaders("Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, " +
	//	"accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Origin, Access-Control-Allow-Credentials")
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	//route.Use(cors.New(config))
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		AllowHeaders: []string{"Content-Type", "Authorization", "Content-Length", "Accept-Encoding", "X-CSRF-Token",
			"accept", "origin", "Cache-Control", "X-Requested-With",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "x-auth-token", "Access-Control-Allow-Headers",
			"token", ""},
		MaxAge: 12 * time.Hour,
	}))

	routes.InitRoutes(route)

	err := route.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
