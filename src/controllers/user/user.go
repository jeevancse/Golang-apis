package user

import (
	"golanguage/src/models"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo"
)

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
	// res := helpers.responseObject("success", 200, true)

	return c.JSON(http.StatusCreated, M{"code": 200, "resp": user, "message": "User created successfully"})

	// return res
	// return helpers.responseObject("success", 200, true)
}
