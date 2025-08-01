package cmd

import (
	config2 "tesodev-korpes/CustomerService/config"
	"tesodev-korpes/CustomerService/internal"
	"tesodev-korpes/pkg"
	//"tesodev-korpes/pkg/middleware"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func BootCustomerService(client *mongo.Client, e *echo.Echo) {
	config := config2.GetCustomerConfig("dev")
	customerCol, err := pkg.GetMongoCollection(client, config.DbConfig.DBName, config.DbConfig.ColName)
	if err != nil {
		panic(err)
	}

	repo := internal.NewRepository(customerCol)
	service := internal.NewService(repo)
	internal.NewHandler(e, service)

	e.Logger.Fatal(e.Start(config.Port))

}
