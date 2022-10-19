package entities

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;type:integer"`
	Username    string    `json:"username" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null"`
	Age         uint      `json:"age" gorm:"type:integer;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}
