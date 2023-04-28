package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEndereco(c *gin.Context) {
	// receber body e o usuario
	var body models.ReqEndereco
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao ler body",
		})

		return
	}

	idCliente, exists := c.Get("idCliente")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "usuario não logado",
		})

		return
	}
	// adicionado endereco
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	_, err := db.Exec(context.Background(), "INSERT INTO endereco(id_cliente, numero, rua, bairro, cidade)VALUES ($1, $2, $3, $4, $5)", idCliente, body.Numero, body.Rua, body.Bairro, body.Cidade)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao adicionar endereço",
		})

		return
	}

	// response

	c.JSON(http.StatusCreated, gin.H{
		"message": "Endereço adicionado",
	})
}

func ReadEnderecos(c *gin.Context) {
	// buscar usuario
	idCliente, exists := c.Get("idCliente")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "usuario não logado",
		})

		return
	}

	// buscar endereços no DB
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), "SELECT * FROM endereco WHERE id_cliente=$1", idCliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buscar no DB",
		})

		return
	}

	var enderecos []models.Endereco

	for rows.Next() {
		var end models.Endereco
		err := rows.Scan(&end.Id, &end.Id_cliente, &end.Numero, &end.Rua, &end.Bairro, &end.Cidade)
		if err != nil {
			continue
		}
		fmt.Println(end)
		enderecos = append(enderecos, end)
	}

	// response

	if enderecos == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "nenhum endereço cadastrado",
		})
	} else {
		c.PureJSON(http.StatusOK, enderecos)
	}
}

func DeleleEndereco(c *gin.Context) {
	// receber param
	idEndereco := c.Param("idEndereco")

	// conferir se o endereço existe
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	var exists bool
	row := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM endereco WHERE id=$1)", idEndereco)
	row.Scan(&exists)

	if exists != true {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "endereço não existe",
		})

		return
	}

	// delete endereço do DB
	fmt.Print(idEndereco)
	_, err := db.Exec(context.Background(), "DELETE FROM endereco WHERE id=$1", idEndereco)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao deletar endereço",
		})

		return
	}

	// response
	c.JSON(http.StatusAccepted, gin.H{
		"message": "endereço deletado",
	})
}
