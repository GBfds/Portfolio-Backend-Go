package controllers

import (
	"Backend-Go/src/initializers"
	"Backend-Go/src/models"
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

	idCliente := c.GetString("idCliente")
	fmt.Println(idCliente)
	// adicionado endereco

	endereco := models.Endereco{Id_cliente: string(idCliente), Numero: body.Numero, Rua: body.Rua, Bairro: body.Bairro, Cidade: body.Cidade}

	result := initializers.DB.Create(&endereco)
	if result.Error != nil {
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
	var enderecos []models.Endereco
	results := initializers.DB.Find(&enderecos, "id_cliente = ?", idCliente)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buscar endereços",
		})

		return
	}

	// response

	if len(enderecos) < 0 {
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
	var endereco models.Endereco
	result := initializers.DB.First(&endereco, "id = ?", idEndereco)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "erro ao buscar endereço",
		})

		return
	}

	if endereco.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "endereco não existe",
		})
	} else {
		// delete endereço do DB
		initializers.DB.Where("id = ?", endereco.ID).Delete(&models.Endereco{})

		// response
		c.JSON(http.StatusAccepted, gin.H{
			"message": "endereço deletado",
		})
	}
}
