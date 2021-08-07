package controllers

import (
	"encoding/json"
	"net/http"
	"store/core/accounts/entities"
	"store/core/accounts/services"
	"store/infra/database/repositories"
)

type AccountController struct {
	Repository repositories.AccountRepository
}

func (ctrl AccountController) Post(rw http.ResponseWriter, r *http.Request) {

	var credentialInput services.Credential

	if err := json.NewDecoder(r.Body).Decode(&credentialInput); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	violations, err := services.CreateCredential(&credentialInput, ctrl.Repository)

	if len(violations) > 0 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(violations)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode("an internal error ocurred when save the credential")
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(credentialInput)
}

func (ctrl AccountController) CreateAccountToCustomer(rw http.ResponseWriter, r *http.Request) {

	var account entities.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	violations, err := services.CreateAccountToCustomer(&account, ctrl.Repository)

	if len(violations) > 0 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(violations)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode("an internal error ocurred when save the credential")
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(account)
}

func (ctrl AccountController) Get() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(ctrl.Repository.GetAllAccounts())
	}
}
