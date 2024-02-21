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

	db := dbConn.Connection()

	sqlDB, err := db.DB()

	//Closing DB connection on panic situation.
	defer sqlDB.Close()

	if err != nil {
		log.Fatalln("Error: on db connection", err)
	}

	if pingErr := sqlDB.Ping(); err != nil {
		log.Fatalln("Error: on db ping", pingErr)
	}

	// Define a new route for GET requests to the "/ping" path.
	// server.GET("/ping", r.Login)
	server.POST("/user/login", r.Login)

	// Start the HTTP server on all available network interfaces, listening on port 8080.
	server.Run() // listen and serve on 0.0.0.0:808(for windows "localhost:8080
	//If we want to use different port we can use
	//server.Run(<port>)
}
