package handlers

import (
	"store/api/controllers"
	"store/infra/database/repositories"

	"github.com/gorilla/mux"
)

var (
	accountsController = controllers.AccountController{Repository: repositories.AccountRepository{}}
)

func ConfigureHandlers(router *mux.Router) {
	router.HandleFunc("/accounts", accountsController.CreateAccountToCustomer).Methods("POST")
	router.HandleFunc("/accounts/credentials", accountsController.Post).Methods("POST")
	router.HandleFunc("/accounts", accountsController.Get()).Methods("GET")
}
