package initializers

import (
	"Backend-Go/src/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.Cliente{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Endereco{})
	DB.AutoMigrate(&models.Produto{})
}
