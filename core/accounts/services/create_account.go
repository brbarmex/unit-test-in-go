package services

import (
	"store/core/accounts/entities"
	"store/core/accounts/interfaces"
	"store/pkg/security"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreateAccountToCustomer(account *entities.Account, repository interfaces.Repository) (violations []string, err error) {

	if violations = account.IsValid(repository.CheckAccountAlreadyExist); len(violations) > 0 {
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
