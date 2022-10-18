package user

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service Service
}

func NewHandler(service Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserRegisterHandler(c *gin.Context) {
	userRequest := UserRegisterRequest{}
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

func (h *UserHandler) UserLoginHandler(c *gin.Context) {
	loginRequest := UserLoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.Login(loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("token", response.Token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UserUpdateHandler(c *gin.Context) {
	user := UserUpdateRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["id"].(float64)
	updatedUser, err := h.service.Update(int(userId), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         updatedUser.ID,
		"email":      updatedUser.Email,
		"username":   updatedUser.Username,
		"age":        updatedUser.Age,
		"updated_at": updatedUser.UpdatedAt,
	})
}
