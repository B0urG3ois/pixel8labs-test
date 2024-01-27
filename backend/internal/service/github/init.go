package github

import (
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
)

func New(
	oauth oauth2.Config,
	config entity.Config,
) *Module {
	return &Module{
		OAuthConfig: oauth,
		Config:      config,
	}
}
