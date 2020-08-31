package model

import "time"

type ChannelEmail struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID  uint
	User    User
	Address string `gorm:"type:varchar(100);unique_index;not null"`
	Count   uint
}
