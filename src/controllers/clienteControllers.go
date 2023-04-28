package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
	"context"
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

	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	_, errorCreate := db.Exec(context.Background(), "INSERT INTO cliente(nome, email, senha, telefone) VALUES ($1, $2, $3, $4)", body.Nome, body.Email, string(hash), body.Telefone)

	if errorCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao criar usuario",
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
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	row := db.QueryRow(context.Background(), "SELECT * FROM cliente WHERE email = $1", body.Email)

	var clt models.Cliente
	errorScan := row.Scan(&clt.Id, &clt.Nome, &clt.Email, &clt.Senha, &clt.Telefone)
	if errorScan != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao receber dados do DB",
		})

		return
	}

	//comparar hash
	err := bcrypt.CompareHashAndPassword([]byte(clt.Senha), []byte(body.Senha))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email/senha incorreto(a)",
		})

		return
	}

	// criar tonken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": clt.Id,
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
