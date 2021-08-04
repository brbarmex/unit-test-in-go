package controllers

import (
	"encoding/json"
	"net/http"
	"store/core/accounts/repositories"
	"store/core/accounts/services"
)

type AccountController struct {
	SendMail   func(string, string)
	Repository repositories.AccountRepository
}

func (ctrl AccountController) PostAccount() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var input services.AccountInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		model, _ := services.CreateNewAccount(input, ctrl.Repository.Insert, ctrl.SendMail)
		json.NewEncoder(rw).Encode(model)
	}
}

func (ctrl AccountController) GetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(ctrl.Repository.SelectAllWithoutPagination())
	}
}
