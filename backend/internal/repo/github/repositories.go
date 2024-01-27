package github

import (
	"context"
	"pixel8labs-test/backend/internal/entity"
)

func (r *Repo) GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error) {
	return r.githubSvc.GetRepositoriesByUser(ctx, reqParams)
}
