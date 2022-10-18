package user

import (
	"time"

	"github.com/mrizalr/mygram/comment"
	"github.com/mrizalr/mygram/photo"
	socialmedia "github.com/mrizalr/mygram/socialMedia"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;type:integer"`
	Username    string    `json:"username" validate:"required" gorm:"type:varchar(255);index:unique;not null"`
	Email       string    `json:"email" validate:"required, email" gorm:"type:varchar(255);index:unique;not null"`
	Password    string    `json:"password" validate:"required, min=6" gorm:"type:varchar(255);not null"`
	Age         uint      `json:"age" validate:"required, numeric, min=8" gorm:"type:integer;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Photos      []photo.Photo
	Comments    []comment.Comment
	SocialMedia []socialmedia.SocialMedia
}
