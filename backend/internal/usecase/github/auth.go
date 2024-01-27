package github

import (
	"context"
	"pixel8labs-test/backend/internal/entity"
	"strings"

	"golang.org/x/oauth2"
)

func (u *Usecase) UserAuth(ctx context.Context, authCode string) (*oauth2.Token, error) {
	return u.GithubRepo.Callback(ctx, authCode)
}

func (u *Usecase) LoginUrl(ctx context.Context) string {
	return u.GithubRepo.UserLogin(ctx)
}

func (u *Usecase) LogoutUrl(ctx context.Context, reqParams entity.OAuthRequest) (bool, error) {
	rawTokens := strings.Split(reqParams.Token, " ")
	if len(rawTokens) < 2 {
		return false, nil
	}

	reqParams.Token = rawTokens[1]
	return u.GithubRepo.UserLogout(ctx, reqParams)
}
