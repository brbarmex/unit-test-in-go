package api

import (
	"net/http"
	"os"
	adapter "store/core/accounts/adapter"
	infra "store/infra/database"

	"github.com/joho/godotenv"
)

func Bootstrapper() {
	db, err := infra.PostgresInitialize()

	if err != nil {
		panic(err)
	}

	godotenv.Load(".env")

	adapter.RouterConfig(db)
	http.ListenAndServe(os.Getenv("port"), nil)
}
