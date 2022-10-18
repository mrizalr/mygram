package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	strToken, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": err.Error(),
		})
		return
	}

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "invalid token",
		})
		return
	}

	c.Set("claims", claims)
	c.Next()
}
