package model

import "time"

type ChannelWeChat struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID uint
	User   User

	MPOpenID string `gorm:"type:varchar(50);unique_index;not null"`
}
