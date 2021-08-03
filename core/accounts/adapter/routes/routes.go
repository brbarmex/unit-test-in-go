package routes

import (
	"encoding/json"
	"net/http"
	"store/core/accounts/repositories"

	"gorm.io/gorm"
)

func GetAccounts(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(rw).Encode(repositories.SelectAllWithoutPagination(db))
	}
}
