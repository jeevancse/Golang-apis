package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`

	FirstName  string `json:"firstname" bson:"firstname"`
	LastName   string `json:"lastname" bson:"lastname"`
	Email      string `json:"email" bson:"email"`
	Mobile     string `json:"mobile" bson:"mobile"`
	Password   string `json:"password" bson:"password"`
	IsVerified bool   `json:"isVerified" bson:"isVerified"`
}
