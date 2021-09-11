package main

import (
	"golanguage/src/config"
	"golanguage/src/models"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

// Create handler create new book.
func Creates(c echo.Context) error {
	type M map[string]interface{}
	var internalError = M{"message": "internal error"}
	if c.FormValue("firstName") == "" {
		return c.JSON(http.StatusBadRequest, M{"code": 400, "success": false, "data": "", "message": "firstName is required"})

	}

	user := models.User{
		FirstName:  c.FormValue("firstName"),
		LastName:   c.FormValue("lastName"),
		Mobile:     c.FormValue("mobile"),
		Email:      c.FormValue("email"),
		Password:   c.FormValue("password"),
		IsVerified: false,
	}

	// To get model's collection, just call to mgm.Coll() method.
	// err := mgm.Coll()
	err := mgm.Coll(&models.User{}).Create(&user)

	if err != nil {
		// Log your error
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	return c.JSON(http.StatusOK, M{"code": 200, "success": true, "message": "Data Inserted successfully.", "data": user})
}

func main() {
	e := echo.New()
	config.Database()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.POST("/add", Creates)
	e.Logger.Fatal(e.Start((":3001")))

}
