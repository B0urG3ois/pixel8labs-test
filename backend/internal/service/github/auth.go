package github

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
	helper "pixel8labs-test/backend/pkg/cryptograph"

	"golang.org/x/oauth2"
)

func (m *Module) Callback(ctx context.Context, authCode string) (*oauth2.Token, error) {
	// Exchange converts an authorization code into a token.
	return m.OAuthConfig.Exchange(ctx, authCode)
}

func (m *Module) Login(ctx context.Context) string {
	// Redirect to GitHub OAuth to Login.
	return m.OAuthConfig.AuthCodeURL(constant.DefaultString, oauth2.AccessTypeOffline)
}

func (m *Module) Logout(ctx context.Context, reqParams entity.OAuthRequest) (bool, error) {
	var (
		result bool
		err    error
		url    = fmt.Sprintf("%s/applications/%s/token", m.Config.BaseUrl, m.Config.ClientID)
	)

	// Create an HTTP client.
	client := &http.Client{}

	// Init body.
	body := map[string]string{"access_token": reqParams.Token}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to marshal body: %v", err)
		return result, err
	}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		return result, err
	}

	// Create request header.
	encClient := helper.Base64Encode(fmt.Sprintf("%s:%s", m.Config.ClientID, m.Config.ClientSecret))
	authHeader := fmt.Sprintf("Basic %s", encClient)
	req.Header.Set(constant.HTTPHeaderAuthorization, authHeader)
	req.Header.Set(constant.HTTPHeaderAcceptContent, constant.DefaultAPIGitHubAcceptContentValue)

	// Do HTTP call to service provider
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to call service provider: %v", err)
		return result, err
	}
	defer res.Body.Close()

	// Validate response status code.
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %v", res.StatusCode)
		return result, errors.New("Request failed")
	}

	return true, nil
}
