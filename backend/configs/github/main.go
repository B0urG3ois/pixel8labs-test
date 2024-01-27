package github

import (
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func InitOAuthConfig(config entity.Config) oauth2.Config {
	return oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user", "username", "repo", "gist"},
	}
}
