package config

import (
	"fmt"
	"os"
)

type HttpServer struct {
	Address string
}

type Config struct {
	Env         string `yaml:"env" env:"env" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoadEnv() {
	configPath := os.Getenv("CONFIG_PATH")

	fmt.Println("Config path value : ", configPath)
}
