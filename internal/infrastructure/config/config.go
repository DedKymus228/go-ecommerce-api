package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string    `yaml:"env" env-default:"dev"`
	DB  DBConfig  `yaml:"db"`
	App AppConfig `yaml:"app"`
}

type DBConfig struct {
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Database string `yaml:"database"`
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
