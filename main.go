package main

import (
	"tesodev-korpes/CustomerService/cmd"
	"tesodev-korpes/pkg"
	"tesodev-korpes/shared/config"

	"github.com/labstack/echo/v4"
)

func main() {

	dbConf := config.GetDBConfig("dev")

	client, err := pkg.GetMongoClient(dbConf.MongoDuration, dbConf.MongoClientURI)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	cmd.BootCustomerService(client, e)

}
