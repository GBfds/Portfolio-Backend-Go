package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUpAdmin(c *gin.Context) {
	// receber body
	var body models.ReqAdimSignUp
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}
	//hash senha
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Senha), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao fazer hash",
		})
	}

	// criando cliente
	admin := models.Admin{Nome: body.Nome, Email: body.Email, Senha: string(hash), Cargo: body.Cargo}
	result := initializers.DB.Create(&admin)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao criar usuario",
		})

		return
	}
	//response
	c.JSON(http.StatusCreated, gin.H{
		"message": "admin criado com sucesso",
	})
}

func LoginAdmin(c *gin.Context) {
	// receber body
	var body models.ReqUserLogin
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}
	// busca no DB
	var admin models.Admin
	initializers.DB.Find(&admin, "email = ?", body.Email)

	//comparar hash
	err := bcrypt.CompareHashAndPassword([]byte(admin.Senha), []byte(body.Senha))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email/senha incorreto(a)",
		})

		return
	}

	// criar tonken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, erro := token.SignedString([]byte(os.Getenv("SECRET")))
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao criar token",
		})

		return
	}

	// response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuthorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"message": tokenString,
	})

	return
}
