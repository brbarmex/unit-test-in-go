package server

import (
	"net/http"
	"os"
	"store/api/handlers"
	"store/api/middlewares"

	"github.com/gorilla/mux"
)

func StartUp() {

	router := mux.NewRouter()
	handlers.ConfigureHandlers(router)
	router.Use(middlewares.SetContentType)
	http.ListenAndServe(os.Getenv("port"), router)
}
