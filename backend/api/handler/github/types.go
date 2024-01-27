package github

import (
	"context"
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
)

type GithubUsecase interface {
	// Auth.
	UserAuth(ctx context.Context, authCode string) (*oauth2.Token, error)
	LoginUrl(ctx context.Context) string
	LogoutUrl(ctx context.Context, reqParams entity.OAuthRequest) (bool, error)

	// User Info.
	GetUserInfo(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error)

	// User Repo.
	GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error)
}

type Handler struct {
	githubUC GithubUsecase
}
