package main

import (
	"database/sql"
	"fmt"
	"rains-backend/controllers"
	"rains-backend/database"
	"rains-backend/middlewares"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST     = "localhost"
	PORT     = 3306
	USER     = "rains_admin"
	PASSWORD = "rainsadmin123"
	DBNAME   = "rains"
)

var DB *sql.DB

func main() {
	dbConnect()

	router := initRouter()
	router.Run(":4000")
}

func dbConnect() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		USER, PASSWORD, HOST, PORT, DBNAME,
	)
	// DB, err := sql.Open("mysql", connString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer DB.Close()
	database.Connect(connString)
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
