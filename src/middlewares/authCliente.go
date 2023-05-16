package middlewares

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthCliente(c *gin.Context) {
	// get cookeis
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// conferindo exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		//
		var cliente models.Cliente
		result := initializers.DB.First(&cliente, "id = ?", claims["sub"])

		if result.Error != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		//atack req
		c.Set("idCliente", cliente.ID)
		//continue
		c.Next()

	} else {
		c.AbortWithError(http.StatusUnauthorized, err)
	}
}
