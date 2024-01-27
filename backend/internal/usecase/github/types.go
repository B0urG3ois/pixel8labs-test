package github

import (
	"context"
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
)

type GithubResource interface {
	// Auth.
	UserLogin(ctx context.Context) string
	UserLogout(ctx context.Context, reqParams entity.OAuthRequest) (bool, error)
	Callback(ctx context.Context, code string) (*oauth2.Token, error)

	// User Info.
	GetUserInfoByUsername(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error)
	GetUserEmails(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.UserEmail, error)
	GetUserByID(ctx context.Context, id int64) (entity.User, error)
	InsertUser(ctx context.Context, newUser entity.User) (int64, error)
	UpdateUserVisitorsByID(ctx context.Context, user entity.User) (int64, error)

	// User Repo.
	GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error)
}

type Usecase struct {
	GithubRepo GithubResource
}
