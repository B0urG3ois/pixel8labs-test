package github

import (
	"context"
	"net/http"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (h *Handler) GetAuthCode(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	result := h.githubUC.LoginUrl(ctx)
	ctx = context.WithValue(ctx, constant.ContextToken, result) //nolint:all

	return result, nil
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	token := r.Header.Get(constant.HTTPHeaderAuthorization)

	return h.githubUC.LogoutUrl(ctx, entity.OAuthRequest{
		Token: token,
	})
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	code := r.URL.Query().Get(constant.ContextCode)
	token, err := h.githubUC.UserAuth(ctx, code)
	if err != nil {
		return token, err
	}

	ctx = context.WithValue(ctx, constant.ContextToken, token.AccessToken) //nolint:all
	return token, nil
}
