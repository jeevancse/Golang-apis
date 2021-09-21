package user

import (
	"context"
	"fmt"
	"golanguage/src/models"
	"log"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
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

func GetAllUsers(c echo.Context) error {
	type M map[string]interface{}
	user := mgm.Coll(&models.User{})
	cursor, err := user.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.Background(), &episodes); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, M{"code": 200, "resp": episodes, "message": "User get successfully"})

}

func DeletesUser(c echo.Context) error {

	id := c.Param("userId")
	fmt.Println(id)
	type M map[string]interface{}
	user := mgm.Coll(&models.User{})
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	result, err := user.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}
	// if result != nil {
	return c.JSON(http.StatusOK, M{"code": 200, "resp": result, "message": "User deleted successfully"})
	// }

	// return nil

}

func GetuserById(c echo.Context) error {
	userId := c.Param("userId")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	type M map[string]interface{}
	user := mgm.Coll(&models.User{})
	var podcast bson.M
	err = user.FindOne(context.Background(), bson.M{"_id": id}).Decode(&podcast)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, M{"code": 200, "resp": podcast, "message": "User get successfully"})
}
