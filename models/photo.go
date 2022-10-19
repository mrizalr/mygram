package models

type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}
