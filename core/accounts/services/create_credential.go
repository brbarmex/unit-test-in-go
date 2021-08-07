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

type Credential struct {
	Id       string `json:"id"`
	Email    string `json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"len=8"`
}

func CreateCredential(credentialInput *Credential, repository interfaces.Repository) (violations []string, er error) {

	if err := validator.Validate(credentialInput); err != nil {
		violations = append(violations, err.Error())
	}

	if _, err := mail.ParseAddress(credentialInput.Email); err != nil {
		violations = append(violations, err.Error())
	}

	if entity := repository.TakeCredentialByEmail(credentialInput.Email); entity.ID > 0 {
		violations = append(violations, "email: Already exist")
	}

	if len(violations) > 0 {
		return violations, nil
	}

	credentialInput.Id = uuid.NewV4().String()

	if err := repository.Create(&entities.Credential{
		ID:        0,
		HashCode:  credentialInput.Id,
		Email:     credentialInput.Email,
		Password:  security.Encrypt(credentialInput.Password),
		Actived:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
