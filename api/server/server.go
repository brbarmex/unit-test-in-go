package server

import (
	"net/http"
	"os"
	"store/api/handlers"
	"store/api/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Rise(db *gorm.DB) {

	router := mux.NewRouter()
	handlers.AccountsRouteConfigure(router, db)

	router.Use(middlewares.SetContentType)
	http.ListenAndServe(os.Getenv("port"), router)
}
