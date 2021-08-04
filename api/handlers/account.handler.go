package handlers

import (
	"encoding/json"
	"net/http"
	"store/core/accounts/repositories"
	"store/core/accounts/services"
	"store/pkg/notifications/email"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AccountsRouteConfigure(router *mux.Router, db *gorm.DB) {

	router.HandleFunc("/accounts", post(db)).Methods("POST")
	router.HandleFunc("/accounts", getAll(db)).Methods("GET")
}

func post(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var input services.AccountInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		model, _ := services.CreateNewAccount(input, db, email.SendMail)
		json.NewEncoder(rw).Encode(model)
	}
}

func getAll(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(repositories.SelectAllWithoutPagination(db))
	}
}
