package model

import "time"

type OAuthGitHub struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID uint
	User   User

	GitHubID uint64 `gorm:"unique_index;not null"`
}
