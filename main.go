package main

import (
	"Backend-Go/src/controllers"
	"Backend-Go/src/initializers"
	"Backend-Go/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.POST("/signup/admin", controllers.SignUpAdmin)
	r.POST("/login/admin", controllers.LoginAdmin)

	r.GET("/enderecos", middlewares.AuthCliente, controllers.ReadEnderecos)
	r.POST("/endereco", middlewares.AuthCliente, controllers.AddEndereco)
	r.DELETE("/endereco/:idEndereco", middlewares.AuthCliente, controllers.DeleleEndereco)

	r.GET("/produtos", controllers.ReadProdutos)
	r.GET("/produto/:idProduto", controllers.ReadUnicoProduto)
	r.POST("/produto", middlewares.AuthAdmin, controllers.CrateProduto)
	r.PUT("/produto/:idProduto", middlewares.AuthAdmin, controllers.UpdateProduto)
	r.DELETE("/produto/:idProduto", middlewares.AuthAdmin, controllers.DeleleProduto)

	r.GET("/", middlewares.AuthCliente, func(c *gin.Context) {
		id, exists := c.Get("idCliente")
		if exists == false {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "",
			})
		}
		c.PureJSON(http.StatusOK, id)
	})

	r.Run()
}
