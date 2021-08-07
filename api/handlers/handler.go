package handlers

import (
	"store/api/controllers"
	"store/infra/database/repositories"

	"github.com/gorilla/mux"
)

var accountsController controllers.AccountController

func ConfigureHandlers(router *mux.Router) {
	loadDependency()
	loadRoutes(router)
}

func loadDependency() {
	accountsController = controllers.AccountController{Repository: repositories.AccountRepository{}}
}

func loadRoutes(r *mux.Router) {
	r.HandleFunc("/accounts/credentials", accountsController.Post).Methods("POST")
	r.HandleFunc("/accounts", accountsController.Get()).Methods("GET")
}
