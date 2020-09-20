package model

import "time"

type OAuthGitHub struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID uint
	User   User

	GitHubID uint `gorm:"unique_index;not null"`
}
