package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (m *Module) GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error) {
	var (
		result []entity.Repository
		url    = fmt.Sprintf("%s/user/repos", m.Config.BaseUrl)
		err    error
	)

	// Set to default user.
	if reqParams.Token == constant.DefaultString {
		url = fmt.Sprintf("%s/users/%s/repos", m.Config.BaseUrl, reqParams.Username)
	}

	// Create an HTTP client.
	client := &http.Client{}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		return result, err
	}

	// Create request header.
	if reqParams.Token != constant.DefaultString {
		req.Header.Set(constant.HTTPHeaderAuthorization, reqParams.Token)
	}
	req.Header.Set(constant.HTTPHeaderAcceptContent, constant.DefaultAPIGitHubAcceptContentValue)

	// Do HTTP call to service provider
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to call service provider: %v", err)
		return result, nil
	}
	defer res.Body.Close()

	// Validate response status code.
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %v", res.StatusCode)
		return result, nil
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return result, nil
	}

	// Parse response body.
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
		return result, err
	}

	return result, nil
}
