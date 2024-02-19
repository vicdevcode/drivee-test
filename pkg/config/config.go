package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env            string        `yaml:"env"             env-required:"true"`
	ContextTimeout time.Duration `yaml:"context_timeout" env-required:"true"`
	Http           HTTPServer    `yaml:"http"            env-required:"true"`
}

type HTTPServer struct {
	Port            string        `yaml:"port"             env-required:"true"`
	Host            string        `yaml:"host"             env-required:"true"`
	ReadTimeout     time.Duration `yaml:"read_timeout"     env-required:"true"`
	WriteTimeout    time.Duration `yaml:"write_timeout"    env-required:"true"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-required:"true"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
