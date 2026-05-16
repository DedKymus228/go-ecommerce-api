package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string    `yaml:"env" env-required:"true"`
	DB  DBConfig  `yaml:"db"  env-required:"true"`
	App AppConfig `yaml:"app" env-required:"true" `
}

type DBConfig struct {
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
}

type AppConfig struct {
	Port        string        `yaml:"serv_port" env-default:"8080"`
	SecretJwt   string        `yaml:"secret_jwt" env-required:"true"`
	RWTimeout   time.Duration `yaml:"rw_timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"120s"`
	TokenTTl    time.Duration `yaml:"token_ttl" env-default:"24h"`
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
