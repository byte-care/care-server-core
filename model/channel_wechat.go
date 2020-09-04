package model

import "time"

type ChannelWeChat struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID uint
	User   User

	UnionID   string `gorm:"type:varchar(50);unique_index;not null"`
	MPOpenID  string `gorm:"type:varchar(50);unique_index"`
	WebOpenID string `gorm:"type:varchar(50);unique_index"`
	APPOpenID string `gorm:"type:varchar(50);unique_index"`
}
