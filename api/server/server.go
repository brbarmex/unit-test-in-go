package server

import (
	"net/http"
	"os"
	"store/api/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Rise(db *gorm.DB) {

	router := mux.NewRouter()
	handlers.AccountsRouteConfigure(router, db)

	router.Use(contentTypeMiddleware)
	http.ListenAndServe(os.Getenv("port"), router)
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
