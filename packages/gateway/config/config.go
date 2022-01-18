package config

import (
	"os"
	"strconv"
)

type Config struct {
	Api *Api
}

func readStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func readIntEnv(key string, defaultValue int) int {
	value, _ := strconv.Atoi(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

func NewConfig() *Config {
	apiPort := readIntEnv(envApiPort, apiDefaultPort)
	return &Config{
		Api: &Api{
			Host: readStringEnv(envApiHost, ""),
			Port: apiPort,
		},
	}
}
