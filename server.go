package main

import (
	dbConn "expense-monster-BE/helper"
	r "expense-monster-BE/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize a new Gin router with the default middleware stack.
	server := gin.Default()

	//Connect to Database
	db := dbConn.Connection()

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error: on db connection", err)
	}

	//Closing DB connection on panic situation.
	defer sqlDB.Close()

	//Create/Update tables in database
	dbConn.AutoMigrate()

	if pingErr := sqlDB.Ping(); err != nil {
		log.Println("Error: on db ping", pingErr)
	}

	server.POST("/user/login", r.Login)

	server.POST("/user/register", r.Register)

	// Start the HTTP server on all available network interfaces, listening on port 8080.
	server.Run() // listen and serve on 0.0.0.0:808(for windows "localhost:8080
	//If we want to use different port we can use
	//server.Run(<port>)
}
