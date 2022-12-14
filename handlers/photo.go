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

	c.JSON(http.StatusCreated, models.ParseToCreatePhotoResponse(photo))
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

	c.JSON(http.StatusOK, photos)
}

func (h *PhotoHandlers) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	photoRequest := models.CreatePhotoRequest{}
	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userID := claims.(jwt.MapClaims)["user_id"].(float64)
	result, err := h.photoService.Update(photoID, int(userID), photoRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ParseToUpdatePhotoResponse(result))
}

func (h *PhotoHandlers) DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	claims, _ := c.Get("claims")
	userID := claims.(jwt.MapClaims)["user_id"].(float64)
	_, err := h.photoService.Delete(photoID, int(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
