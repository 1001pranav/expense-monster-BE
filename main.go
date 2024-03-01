package main

import (
	"expense-monster-BE/handler"
	dbConn "expense-monster-BE/helper"

	// "expense-monster-BE/middleware"
	"expense-monster-BE/docs"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Expense Monster BE
// @version 1.0
// @description This is a sample server for Expense Monster BE API.
// @termsOfService https://example.com/terms/
// @contact npranavr@gmail.com
// @license MIT
// @host localhost:8080
// @BasePath /

func main() {
	docs.SwaggerInfo.Title = "Expense Monster"
	docs.SwaggerInfo.Description = "Backend APIs for Expense Monster "

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

	userRoute := server.Group("/user/")
	{
		//User Login
		//@Tags Login V1
		//@Summary User Login
		//@Description User Login
		//@Accept json
		//@Produce json
		//@Success 200 {object} handler.LoginResponse
		//@Failure 400 {object} handler.ErrorResponse
		//@Failure 401 {object} handler.ErrorResponse
		//@Failure 500 {object} handler.ErrorResponse
		//@Router /user/login
		userRoute.POST("login", handler.Login)
		// Serve Swagger UI
		userRoute.POST("register", handler.Register)

	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// server.Use(middleware.ValidateAccessToken())
	// Start the HTTP server on all available network interfaces, listening on port 8080.
	server.Run() // listen and serve on 0.0.0.0:808(for windows "localhost:8080
	//If we want to use different port we can use
	//server.Run(<port>)
}
