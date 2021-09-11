package config

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database() {
	_ = mgm.SetDefaultConfig(nil, "check_echo", options.Client().ApplyURI("mongodb://localhost:27017/myDatabase"))
	fmt.Println("Database connected successfully.")

}
