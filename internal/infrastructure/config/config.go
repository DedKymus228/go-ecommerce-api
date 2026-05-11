package config

import (
	"log"

	"e-commerce-api/pkg/postgre"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string           `yaml:"env" env-default:"dev"`
	DB  postgre.DBConfig `yaml:"db"`
	App AppConfig        `yaml:"app"`
}

type AppConfig struct {
	Port string `yaml:"serv_port" env-default:"8080"`
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
