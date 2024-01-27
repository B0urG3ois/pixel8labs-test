package github

import (
	"net/http"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	token := r.Header.Get(constant.HTTPHeaderAuthorization)

	return h.githubUC.GetUserInfo(ctx, entity.OAuthRequest{
		Token: token,
	})
}
