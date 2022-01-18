package main

import (
	"github.com/Ibrahimsyah/temasi-go/packages/gateway/config"
	"github.com/Ibrahimsyah/temasi-go/packages/gateway/routers"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.NewConfig()
	server := echo.New()

	routers.NewCommonRouter(server)

	if err := server.Start(config.Api.GetPort()); err != nil {
		server.Logger.Fatal(err)
	}
}
