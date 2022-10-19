package entities

import "time"

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primaryKey;type:integer"`
	Name           string    `json:"name" validate:"required" gorm:"type:varchar(255);not null"`
	SocialMediaURL string    `json:"social_media_url" validate:"required" gorm:"type:varchar(255);not null"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
