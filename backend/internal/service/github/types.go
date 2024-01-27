package github

import (
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
)

type Module struct {
	OAuthConfig oauth2.Config
	Config      entity.Config
}
