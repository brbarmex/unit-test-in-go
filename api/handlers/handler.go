package handlers

import (
	"store/api/controllers"
	"store/core/accounts/repositories"
	databases "store/infra/database"
	"store/pkg/notifications/email"

	"github.com/gorilla/mux"
)

var accountsController controllers.AccountController

func ConfigureHandlers(router *mux.Router) {
	loadDependency()
	loadRoutes(router)
}

func loadDependency() {
	db := databases.Load()
	accountsController = controllers.AccountController{SendMail: email.SendMail, Repository: repositories.AccountRepository{Database: db}}
}

func loadRoutes(r *mux.Router) {
	r.HandleFunc("/accounts", accountsController.PostAccount()).Methods("POST")
	r.HandleFunc("/accounts", accountsController.GetAll()).Methods("GET")
}
