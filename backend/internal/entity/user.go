package entity

import "time"

type User struct {
	ID           int64     `json:"id"`
	AvatarURL    string    `json:"avatar_url"`
	Name         string    `json:"name"`
	Username     string    `json:"login"`
	Email        string    `json:"email"`
	Bio          string    `json:"bio"`
	Followers    int64     `json:"followers"`
	Following    int64     `json:"following"`
	FollowersUrl string    `json:"followers_url"`
	FollowingUrl string    `json:"following_url"`
	Visitors     int64     `json:"visitors"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserEmail struct {
	Email      string `json:"email"`
	IsPrimary  bool   `json:"primary"`
	Visibility string `json:"visibility"`
}
