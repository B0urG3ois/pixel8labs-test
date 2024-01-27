package github

import (
	"context"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (u *Usecase) GetUserInfo(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error) {
	var (
		result entity.User
		err    error
	)

	// Get User Info.
	result, err = u.getUserDetails(ctx, entity.OAuthRequest{
		Username: constant.DefaultUser,
		Token:    reqParams.Token,
	})
	if err != nil {
		return result, err
	}

	// Insert New User into DB.
	_, err = u.GithubRepo.InsertUser(ctx, result)
	if err != nil {
		return result, err
	}

	// Count visitors.
	result.Visitors, err = u.GithubRepo.UpdateUserVisitorsByID(ctx, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *Usecase) getUserDetails(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error) {
	var (
		user       entity.User
		userEmails []entity.UserEmail
		err        error
	)

	// Get User.
	user, err = u.GithubRepo.GetUserInfoByUsername(ctx, reqParams)
	if err != nil {
		return user, err
	}

	user, err = u.GithubRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return user, err
	}

	// Get User Emails.
	userEmails, err = u.GithubRepo.GetUserEmails(ctx, reqParams)
	if err != nil {
		return user, err
	}

	// Get User's primary email
	if len(userEmails) > 0 {
		user.Email = userEmails[0].Email
	}

	return user, nil
}
