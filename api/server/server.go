package server

import (
	"net/http"
	"os"
	"store/api/handlers"
	"store/api/middlewares"
	databases "store/infra/database"

	"github.com/gorilla/mux"
)

func StartUp() {
	databases.Load()
	router := mux.NewRouter()
	handlers.ConfigureHandlers(router)
	router.Use(middlewares.SetContentType)
	http.ListenAndServe(os.Getenv("port"), router)
}
