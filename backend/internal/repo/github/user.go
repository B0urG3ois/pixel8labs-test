package github

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"pixel8labs-test/backend/internal/constant"
	"pixel8labs-test/backend/internal/entity"
)

func (r *Repo) GetUserInfoByUsername(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error) {
	return r.githubSvc.GetUserInfo(ctx, reqParams)
}

func (r *Repo) GetUserEmails(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.UserEmail, error) {
	return r.githubSvc.GetUserEmails(ctx, reqParams)
}

func (r *Repo) GetUserByID(ctx context.Context, id int64) (entity.User, error) {
	var (
		result entity.User
		uTable userTable
		err    error
	)

	row := r.db.QueryRowContext(ctx, querySelectByID, id)
	if err = row.Scan(
		&uTable.ID,
		&uTable.AvatarURL,
		&uTable.Name,
		&uTable.Username,
		&uTable.Email,
		&uTable.Bio,
		&uTable.Followers,
		&uTable.Following,
		&uTable.FollowersUrl,
		&uTable.FollowingUrl,
		&uTable.Visitors,
		&uTable.CreatedAt,
		&uTable.UpdatedAt,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return result, err
	}

	return uTable.ToEntity(), nil
}

func (r *Repo) InsertUser(ctx context.Context, newUser entity.User) (int64, error) {
	var (
		result int64
		err    error
	)

	existingUser, err := r.GetUserByID(ctx, newUser.ID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return result, err
	}
	if existingUser.ID != constant.DefaultInt64 {
		return existingUser.ID, nil
	}

	data := r.toBase(newUser)
	row := r.db.QueryRowContext(ctx, queryInsertFromBaseRepo,
		data.ID,
		data.AvatarURL,
		data.Name,
		data.Username,
		data.Email,
		data.Bio,
		data.Followers,
		data.Following,
		data.FollowersUrl,
		data.FollowingUrl,
		data.Visitors,
		data.CreatedAt,
		data.UpdatedAt,
	)
	if err = row.Scan(&result); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return result, err
	}

	return result, nil
}

func (r *Repo) UpdateUserVisitorsByID(ctx context.Context, user entity.User) (int64, error) {
	var (
		result int64
		err    error
	)

	user.Visitors = user.Visitors + 1
	data := r.toBase(user)
	_, err = r.db.ExecContext(ctx, queryUpdateVisitorsByID,
		data.Visitors,
		data.ID,
	)
	if err != nil {
		return result, err
	}

	return user.Visitors, nil
}
