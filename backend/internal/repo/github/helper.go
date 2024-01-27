package github

import (
	"database/sql"
	"pixel8labs-test/backend/internal/entity"
)

func (r *Repo) toBase(args entity.User) userTable {
	return userTable{
		ID:        args.ID,
		AvatarURL: args.AvatarURL,
		Name:      args.Name,
		Username: sql.NullString{
			String: args.Username,
			Valid:  true,
		},
		Email:        args.Email,
		Followers:    args.Followers,
		Following:    args.Following,
		FollowersUrl: args.FollowersUrl,
		FollowingUrl: args.FollowingUrl,
		Visitors:     args.Visitors,
		CreatedAt:    args.CreatedAt,
		UpdatedAt:    args.UpdatedAt,
	}
}
