package config

import (
	"e-commerce-api/internal/server"
	"log"

	"e-commerce-api/pkg/postgre"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string           `yaml:"env" env-required:"true"`
	DB  postgre.DBConfig `yaml:"db"  env-required:"true"`
	App server.AppConfig `yaml:"app" env-required:"true" `
}

func GetConfig(path string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatal("error loading config file: " + err.Error())
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("error loading env variables: " + err.Error())
	}

	return &cfg

}
