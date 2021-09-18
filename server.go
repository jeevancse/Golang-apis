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
	router.GET("/get-user", user.GetAllUsers)
	// router.DELETE("/delete", user.DeletesUser)

	e.Logger.Fatal(e.Start((":3001")))

}

// func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set("x-version", "Test/v1.0")
// 		return next(c)
// 	}
// }
