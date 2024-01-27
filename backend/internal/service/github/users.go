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

func (m *Module) GetUserInfo(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error) {
	var (
		result entity.User
		url    = fmt.Sprintf("%s/user", m.Config.BaseUrl)
		err    error
	)

	// Set to default user.
	if reqParams.Token == constant.DefaultString {
		url = fmt.Sprintf("%s/users/%s", m.Config.BaseUrl, reqParams.Username)
	}

	// Create an HTTP client.
	client := &http.Client{}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
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
		log.Printf("Failed to call service provider: %v", err)
		return result, nil
	}
	defer res.Body.Close()

	// Validate response status code.
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %v", res.StatusCode)
		return result, nil
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return result, nil
	}

	// Parse response body.
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Printf("Error unmarshaling response body: %v", err)
		return result, err
	}

	return result, nil
}

func (m *Module) GetUserEmails(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.UserEmail, error) {
	var (
		result []entity.UserEmail
		url    = fmt.Sprintf("%s/user/emails", m.Config.BaseUrl)
		err    error
	)
	result = append(result, entity.UserEmail{
		Email: constant.DefaultEmail,
	})

	// Set to default user.
	if reqParams.Token == constant.DefaultString {
		return result, nil
	}

	// Create an HTTP client.
	client := &http.Client{}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
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
		log.Printf("Failed to call service provider: %v", err)
		return result, nil
	}
	defer res.Body.Close()

	// Validate response status code.
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %v", res.StatusCode)
		return result, nil
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return result, nil
	}

	// Parse response body.
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Printf("Error unmarshaling response body: %v", err)
		return result, err
	}

	return result, nil
}
