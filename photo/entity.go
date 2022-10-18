package photo

import (
	"time"

	"github.com/mrizalr/mygram/comment"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:integer"`
	Title     string    `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" validate:"required" gorm:"type:varchar(255);not null"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []comment.Comment
}
