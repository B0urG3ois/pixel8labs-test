package github

import (
	"context"
	"pixel8labs-test/backend/internal/entity"

	"golang.org/x/oauth2"
)

func (r *Repo) UserLogin(ctx context.Context) string {
	return r.githubSvc.Login(ctx)
}

func (r *Repo) UserLogout(ctx context.Context, reqParams entity.OAuthRequest) (bool, error) {
	return r.githubSvc.Logout(ctx, reqParams)
}

func (r *Repo) Callback(ctx context.Context, code string) (*oauth2.Token, error) {
	return r.githubSvc.Callback(ctx, code)
}
