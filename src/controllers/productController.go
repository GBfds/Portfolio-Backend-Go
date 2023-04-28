package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrateProduto(c *gin.Context) {
	// receber body
	var body models.ReqProduto
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}

	//conferir se há produto com mesmo nome
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	var exists bool
	row := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM produto WHERE nome = $1 AND tamanho=$2)", body.Nome, body.Tamanho)
	row.Scan(&exists)

	if exists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "produto já existe",
		})

		return
	}

	// adicionar produto no DB
	_, err := db.Exec(context.Background(), "INSERT INTO produto(nome, tamanho, preco) VALUES ($1, $2, $3)", body.Nome, body.Tamanho, body.Preco)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha as adicionar produto",
		})

		return
	}

	// response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Produto adicionado com sucesso",
	})
}

func ReadProdutos(c *gin.Context) {
	// buscar produtos no DB
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), "SELECT * FROM produto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buscar produtos",
		})

		return
	}

	var produtos []models.Produto
	for rows.Next() {
		var clt models.Produto
		err := rows.Scan(&clt.Id, &clt.Nome, &clt.Tamanho, &clt.Preco)
		if err != nil {
			continue
		}

		produtos = append(produtos, clt)
	}

	//response
	if produtos == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "não há produtos cadastrados",
		})
	} else {
		c.PureJSON(http.StatusAccepted, produtos)
	}
}

func ReadUnicoProduto(c *gin.Context) {
	// receber params
	idParam := c.Param("idProduto")

	//conferir se o produto existe
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	var exists bool
	existsRow := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM produto WHERE id = $1)", idParam)
	existsRow.Scan(&exists)

	if exists != true {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	}

	// buscar produto no DB
	row := db.QueryRow(context.Background(), "SELECT * FROM produto WHERE id=$1", idParam)
	var produto models.Produto

	errScan := row.Scan(&produto.Id, &produto.Nome, &produto.Tamanho, &produto.Preco)
	if errScan != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao receber dados do DB",
		})

		return
	}

	//response
	c.PureJSON(http.StatusFound, produto)
}

func UpdateProduto(c *gin.Context) {
	//receber body e param
	var body models.Produto
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}

	body.Id = c.Param("idProduto")

	// conferir se o produto existe
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	var exists bool
	existsRow := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM produto WHERE id = $1)", body.Id)
	existsRow.Scan(&exists)

	if exists != true {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	}

	//update produto
	_, err := db.Exec(context.Background(), "UPDATE produto SET nome = $1, tamanho = $2, preco = $3 WHERE id = $4", body.Nome, body.Tamanho, body.Preco, body.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao editar produto",
		})

		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto atualizado com sucesso",
	})
}

func DeleleProduto(c *gin.Context) {
	// receber param
	idParam := c.Param("idProduto")

	// conferir se o produto existe
	db := initializers.ConnectToDB()
	defer db.Close(context.Background())

	var exists bool
	existsRow := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM produto WHERE id = $1)", idParam)
	existsRow.Scan(&exists)

	if exists != true {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	}

	// Deletar produto do DB
	_, err := db.Exec(context.Background(), "DELETE FROM produto WHERE id=$1", idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao deletar Produto",
		})

		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto deletado",
	})
}
