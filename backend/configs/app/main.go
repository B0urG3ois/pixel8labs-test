package app

import (
	"os"
	"pixel8labs-test/backend/internal/entity"
)

func InitAppConfig() entity.Config {
	return entity.Config{
		GithubConfig: entity.GithubConfig{
			BaseUrl:      os.Getenv("GITHUB_BASE_URL"),
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		},
		ApplicationConfig: entity.ApplicationConfig{
			Port: os.Getenv("PORT"),
		},
		DatabaseConfig: entity.DatabaseConfig{
			Path: os.Getenv("DB_PATH"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
