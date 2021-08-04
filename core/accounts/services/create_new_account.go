package services

import (
	"store/core/accounts/entities"

	uuid "github.com/satori/go.uuid"
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

func CreateNewAccount(input AccountInput, saveRegister func(entity *entities.Account), sendEmail func(msg string, to string)) (AccountOutput, Notification) {

	account := entities.Account{
		HashCode: uuid.NewV4().String(),
		UserName: input.Name,
		Email:    input.Email,
		Password: input.Password,
		Actived:  true,
	}

	saveRegister(&account)

	go sendEmail("", input.Email)

	return AccountOutput{Id: account.HashCode}, Notification{}
}
