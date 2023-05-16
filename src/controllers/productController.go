package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
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

	// adicionar produto no DB

	produto := models.Produto{Nome: body.Nome, Tamanho: body.Tamanho, Preco: body.Preco}
	result := initializers.DB.Create(&produto)

	if result.Error != nil {
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
	var produtos []models.Produto
	results := initializers.DB.Find(&produtos)

	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buscar produtos",
		})
	}

	//response
	if len(produtos) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "não há produtos cadastrados",
		})
	} else {
		c.PureJSON(http.StatusAccepted, produtos)
	}
}

func ReadUnicoProduto(c *gin.Context) {
	// receber params
	idProduto := c.Param("idProduto")

	// buscar produto no DB
	var produto models.Produto
	result := initializers.DB.First(&produto, "id = ?", idProduto)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buacar produto",
		})

		return
	}

	if produto.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	}

	//response
	c.PureJSON(http.StatusFound, produto)
}

func UpdateProduto(c *gin.Context) {
	//receber body e param
	var body models.ReqProduto
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao ler o body",
		})

		return
	}

	idProduto := c.Param("idProduto")

	// conferir se o produto existe
	var produto models.Produto
	result := initializers.DB.First(&produto, "id = ?", idProduto)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buacar produto",
		})

		return
	}

	if produto.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	}

	//update produto
	initializers.DB.Where("id = ?", produto.ID).Updates(&models.Produto{Nome: body.Nome, Tamanho: body.Tamanho, Preco: body.Preco})

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto atualizado com sucesso",
	})
}

func DeleleProduto(c *gin.Context) {
	// receber param
	idProduto := c.Param("idProduto")

	// conferir se o produto existe
	var produto models.Produto
	resultSelect := initializers.DB.First(&produto, "id = ?", idProduto)
	if resultSelect.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "falha ao deletar Produto",
		})

		return
	}

	if produto.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produto não existe",
		})

		return
	} else {
		// Deletar produto do DB
		initializers.DB.Delete(&models.Produto{}, "id = ?", produto.ID)

		//response
		c.JSON(http.StatusOK, gin.H{
			"message": "Produto deletado",
		})
	}

}
