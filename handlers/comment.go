package handlers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/services"
)

type CommentHandlers struct {
	commentService services.CommentService
}

func NewCommentHandlers(commentService services.CommentService) *CommentHandlers {
	return &CommentHandlers{
		commentService: commentService,
	}
}

func (h *CommentHandlers) CreateComment(c *gin.Context) {
	createCommentRequest := models.CreateCommentRequest{}
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(float64)

	comment, err := h.commentService.Create(int(userId), createCommentRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (h *CommentHandlers) GetAllComment(c *gin.Context) {
	c.MustGet("claims")

	comments, err := h.commentService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandlers) UpdateComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	commentRequest := models.UpdateCommentRequest{}
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userID := claims.(jwt.MapClaims)["user_id"].(float64)

	comment, err := h.commentService.Update(commentID, int(userID), commentRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandlers) DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, _ := c.Get("claims")
	userID := claims.(jwt.MapClaims)["user_id"].(float64)

	comment, err := h.commentService.Delete(commentID, int(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}
