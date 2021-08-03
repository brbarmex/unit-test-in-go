package adapter

import (
	"net/http"
	"store/core/accounts/adapter/routes"

	"gorm.io/gorm"
)

func RouterConfig(db *gorm.DB) {
	http.Handle("/api/accounts", routes.GetAccounts(db))
}
