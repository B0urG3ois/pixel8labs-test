package github

import (
	"context"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (u *Usecase) GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error) {
	var (
		result []entity.Repository
		err    error
	)

	result, err = u.GithubRepo.GetRepositoriesByUser(ctx, entity.OAuthRequest{
		Username: constant.DefaultUser,
		Token:    reqParams.Token,
	})
	if err != nil {
		return result, err
	}

	if len(result) > 6 {
		result = result[:6]
	}

	return result, nil
}
