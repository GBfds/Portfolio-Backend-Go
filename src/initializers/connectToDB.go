package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// postgres://furgbznv:4zNTgUsSEeX0usWhh8BLE_6S6kmhJB0k@tuffi.db.elephantsql.com/furgbznv
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("faldo conectar o DB")
	}
}
