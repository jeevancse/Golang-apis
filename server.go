package main

import (
	"golanguage/src/config"
	"golanguage/src/controllers/user"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Database()
	router := e.Group("/admin")
	router.POST("/create-user", user.Creates)
	e.Logger.Fatal(e.Start((":3001")))

}
