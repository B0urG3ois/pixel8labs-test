package github

import (
	"context"
	"database/sql"
	"pixel8labs-test/backend/internal/entity"
	"time"

	"golang.org/x/oauth2"
)

type githubResource interface {
	// Auth.
	Callback(ctx context.Context, authCode string) (*oauth2.Token, error)
	Login(ctx context.Context) string
	Logout(ctx context.Context, reqParams entity.OAuthRequest) (bool, error)

	// User Info.
	GetUserInfo(ctx context.Context, reqParams entity.OAuthRequest) (entity.User, error)
	GetUserEmails(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.UserEmail, error)

	// User Repo.
	GetRepositoriesByUser(ctx context.Context, reqParams entity.OAuthRequest) ([]entity.Repository, error)
}

type dbTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repo struct {
	githubSvc githubResource
	db        dbTX
}

type userTable struct {
	ID           int64
	AvatarURL    string
	Name         string
	Username     sql.NullString
	Email        string
	Followers    int64
	Following    int64
	FollowersUrl string
	FollowingUrl string
	Visitors     int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *userTable) ToEntity() entity.User {
	return entity.User{
		ID:           u.ID,
		AvatarURL:    u.AvatarURL,
		Name:         u.Name,
		Username:     u.Username.String,
		Email:        u.Email,
		Followers:    u.Followers,
		Following:    u.Following,
		FollowersUrl: u.FollowersUrl,
		FollowingUrl: u.FollowingUrl,
		Visitors:     u.Visitors,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
