package controllers

import (
	"encoding/json"
	"net/http"
	"store/core/accounts/services"
	"store/infra/database/repositories"
)

type AccountController struct {
	Repository repositories.AccountRepository
}

func (ctrl AccountController) Post() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var credentialInput services.Credential

		if err := json.NewDecoder(r.Body).Decode(&credentialInput); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if errors := services.CreateCredential(&credentialInput, ctrl.Repository); len(errors) > 0 {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(errors)
			return
		}

		json.NewEncoder(rw).Encode(credentialInput)
	}
}

func (ctrl AccountController) Get() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("oi")
	}
}
