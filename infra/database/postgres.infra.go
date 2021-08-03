package infra

import (
	"fmt"
	"os"
	"store/core/accounts/entities"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresInitialize() (db *gorm.DB, err error) {

	godotenv.Load(".env")

	if db, err = gorm.Open(postgres.Open(os.Getenv("db_connection")), &gorm.Config{}); err != nil {
		fmt.Println("Failed during connection database.")
		panic(err)
	}

	if enableMigrations, err := strconv.ParseBool(os.Getenv("enable_migrations")); err == nil && enableMigrations {
		if err := db.AutoMigrate(&entities.Account{}); err != nil {
			fmt.Println("Failed during migration executation.")
			panic(err)
		}
		fmt.Println("Migration executed has success.")
	}

	return db, nil
}
