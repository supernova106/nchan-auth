package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port      string
	Oauth2Url string
	GinEnv    string
}

func Load(envFile string) (*Config, error) {
	if envFile == "" {
		envFile = ".env"
	}

	fmt.Printf("Loading %s file\n", envFile)

	godotenv.Load(envFile)

	oauth2Url := os.Getenv("OAUTH2_URL")
	if oauth2Url == "" {
		return nil, fmt.Errorf("Missing OAUTH2_URL")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	ginEnv := os.Getenv("GIN_ENV")
	if ginEnv == "" {
		ginEnv = "development"
	}

	return &Config{
		port,
		oauth2Url,
		ginEnv,
	}, nil
}
