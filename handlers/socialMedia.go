package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/services"
)

type SocialMediaHandlers struct {
	socmedService services.SocmedService
}

func NewSocialMediaHandlers(socmedService services.SocmedService) *SocialMediaHandlers {
	return &SocialMediaHandlers{
		socmedService: socmedService,
	}
}

func (h *SocialMediaHandlers) AddSocialMedia(c *gin.Context) {
	socmedRequest := models.AddSocialMediaRequest{}
	if err := c.ShouldBindJSON(&socmedRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(float64)

	socialMedia, err := h.socmedService.Add(int(userId), socmedRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}
