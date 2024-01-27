package entity

import "time"

type Repository struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	Topics      []string  `json:"topics"`
	Language    string    `json:"language"`
	Url         string    `json:"url"`
	IsPrivate   bool      `json:"private"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
