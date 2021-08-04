package services

import (
	"store/core/accounts/entities"
	"store/core/accounts/repositories"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountOutput struct {
	Id string `json:"id"`
}

type Notification struct {
	messages []string
}

func CreateNewAccount(input AccountInput, db *gorm.DB, sendEmail func(msg string, to string)) (AccountOutput, Notification) {

	account := entities.Account{
		HashCode: uuid.NewV4().String(),
		UserName: input.Name,
		Email:    input.Email,
		Password: input.Password,
		Actived:  true,
	}

	repositories.Insert(&account, db)

	go sendEmail("", input.Email)

	return AccountOutput{Id: account.HashCode}, Notification{}
}
