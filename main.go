package main

import (
	"rains-backend/controllers"
	"rains-backend/database"
	"rains-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnect()

	router := initRouter()
	router.Run(":4000")
}

func dbConnect() {
	database.Connect("rainsadmin:rainsadmin123@tcp(localhost:3306)/rains?parseTime=true")
	database.Migrate()
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/login", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
