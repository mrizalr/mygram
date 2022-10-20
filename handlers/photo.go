package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/services"
)

type PhotoHandlers struct {
	photoService services.PhotoService
}

func NewPhotoHandlers(photoService services.PhotoService) *PhotoHandlers {
	return &PhotoHandlers{
		photoService: photoService,
	}
}

func (h *PhotoHandlers) UploadPhoto(c *gin.Context) {
	createPhotoRequest := models.CreatePhotoRequest{}
	if err := c.ShouldBindJSON(&createPhotoRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(float64)

	photo, err := h.photoService.Create(int(userId), createPhotoRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (h *PhotoHandlers) GetAllPhotos(c *gin.Context) {
	photos, err := h.photoService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	photosResponse := []models.GetPhotoResponse{}
	for _, photo := range photos {
		photosResponse = append(photosResponse, models.ParseToGetPhotoResponse(photo, photo.User))
	}

	c.JSON(http.StatusOK, photosResponse)
}
