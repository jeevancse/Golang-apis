package main

import (
	"golanguage/src/config"
	"golanguage/src/controllers/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	config.Database()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router := e.Group("/admin")
	router.POST("/create-user", user.Creates)
	router.GET("/get-user", user.GetAllUsers)
	router.DELETE("/delete-user/:userId", user.DeletesUser)
	router.GET("/get-user-by-id/:userId", user.GetuserById)
	e.Logger.Fatal(e.Start((":3001")))

}
