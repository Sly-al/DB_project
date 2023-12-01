package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode" env-default:"disable"`
}

func MustLoad() *Config {
	var cfg Config
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Config error: %s", err)
	}
	configPath := os.Getenv("Config_Path")
	if configPath == "" {
		log.Fatal("Config path is not set")
	}
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
