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

func SignUp(c *gin.Context) {
	// receber body
	var body models.ReqClienteSignUp
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

		return
	}

	// criando cliente
	cliente := models.Cliente{Nome: body.Nome, Email: body.Email, Senha: string(hash), Telefone: body.Telefone}
	result := initializers.DB.Create(&cliente)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao criar usuario",
		})

		return
	}

	//response
	c.JSON(http.StatusCreated, gin.H{
		"message": "usuário criado com sucesso",
	})
}

func Login(c *gin.Context) {
	// receber body
	var body models.ReqUserLogin
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}

	// confirir no DB
	var cliente models.Cliente
	initializers.DB.Find(&cliente, "email = ?", body.Email)

	if cliente.ID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email/senha invalido",
		})

		return
	}

	//comparar hash
	err := bcrypt.CompareHashAndPassword([]byte(cliente.Senha), []byte(body.Senha))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email/senha incorreto(a)",
		})

		return
	}

	// criar tonken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": cliente.ID,
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
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"message": tokenString,
	})

	return
}
