package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	UserRegisterHandler(c *gin.Context)
}

type userHandler struct {
	service Service
}

func NewHandler(service Service) *userHandler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) UserRegisterHandler(c *gin.Context) {
	userRequest := UserRequest{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	newUser, err := h.service.Register(userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"email":    newUser.Email,
		"username": newUser.Username,
		"age":      newUser.Age,
	})
}
