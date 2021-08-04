package api

import (
	"store/api/server"
	databases "store/infra/database"
)

func Execute() {
	if db, err := databases.Load(); err == nil {
		server.Rise(db)
	} else {
		panic(err)
	}
}
