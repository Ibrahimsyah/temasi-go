package config

import "fmt"

var (
	envApiHost     = "APP_HOST"
	envApiPort     = "APP_PORT"
	apiDefaultHost = "localhost"
	apiDefaultPort = 5000
)

type Api struct {
	Host string
	Port int
}

func (api Api) GetAddress() string {
	return fmt.Sprintf("%s:%d", api.Host, api.Port)
}

func (api Api) GetPort() string {
	return fmt.Sprintf(":%d", api.Port)
}
