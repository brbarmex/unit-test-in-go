package services

import (
	"net/mail"
	"store/core/accounts/entities"
	"store/core/accounts/interfaces"
	"store/pkg/security"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/validator.v2"
)

func CreateAccountToCustomer(account *entities.Account, repository interfaces.Repository) (violations []string, err error) {

	if error := validator.Validate(account); error != nil {
		violations = append(violations, error.Error())
	}

	if _, error := mail.ParseAddress(account.Email); error != nil {
		violations = append(violations, error.Error())
	}

	exist := repository.CheckAccountAlreadyExist(account.Email, account.Identification)

	if exist {
		violations = append(violations, "email or identification already exist")
	}

	if len(violations) > 0 {
		return violations, nil
	}

	account.HashCode = uuid.NewV4().String()
	account.Password = security.Encrypt(account.Password)
	account.CreatedAt = time.Time{}
	account.UpdatedAt = time.Time{}

	if err = repository.SaveAccount(account); err != nil {
		return nil, err
	}

	return nil, nil
}
