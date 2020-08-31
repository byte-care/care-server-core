package model

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"not null"`

	AccessKey string `gorm:"size:44;unique;not null"`
	SecretKey string `gorm:"size:44;not null"`
}
