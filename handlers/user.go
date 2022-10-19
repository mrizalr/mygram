package handlers

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserRegisterHandler(c *gin.Context) {
	userRequest := models.UserRegisterRequest{}
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
	loginRequest := models.UserLoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	user, err := h.service.Login(loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": user.ID})
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("token", signedToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}

func (h *UserHandler) UserUpdateHandler(c *gin.Context) {
	user := models.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(float64)
	updatedUser, err := h.service.Update(int(userId), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         updatedUser.ID,
		"email":      updatedUser.Email,
		"username":   updatedUser.Username,
		"age":        updatedUser.Age,
		"updated_at": updatedUser.UpdatedAt,
	})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(float64)
	_, err := h.service.Delete(int(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("token", "", 0, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
