package api

import (
	"net/http"
	"os"
	adapter "store/core/accounts/adapter"
	infra "store/infra/database"
)

func Bootstrapper() {
	db, err := infra.PostgresInitialize()

	if err != nil {
		panic(err)
	}

	adapter.RouterConfig(db)
	http.ListenAndServe(os.Getenv("port"), nil)
}
