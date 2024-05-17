package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type (
	// Config -.
	Config struct {
		App     `yaml:"app"`
		Handler `yaml:"handler"`
		Redis   `yaml:"redis"`
		MySQL   `yaml:"mysql"`
	}

	App struct {
		Name    string `env:"APP_NAME" env-default:"Todo List App"`
		Version string `env:"APP_VERSION" env-default:"v1.1.1"`
	}

	Handler struct {
		DBClient string `env:"DB_CLIENT" env-default:"redis"`
	}

	Redis struct {
		Password string `env:"REDIS_PASSWORD" env-default:""`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
		DB       int    `env:"REDIS_DATABASES" env-default:"1"`
	}

	MySQL struct {
		Host     string `env:"MYSQL_HOST" env-default:"localhost"`
		Port     string `env:"MYSQL_PORT" env-default:"3306"`
		User     string `env:"MYSQL_USER" env-default:"root"`
		Password string `env:"MYSQL_PASSWORD" env-default:""`
		DataBase string `env:"MYSQL_DATABASE" env-default:"test"`
	}
)

var MainConfig = &Config{}

func init() {
	err := cleanenv.ReadEnv(MainConfig)
	if err != nil {
		log.Fatal(err)
	}

	if MainConfig.Handler.DBClient != "redis" && MainConfig.Handler.DBClient != "mysql" {
		log.Fatal("DB_CLIENT должен быть либо 'redis', либо 'mysql'")
	}
}
