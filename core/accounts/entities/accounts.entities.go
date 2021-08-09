package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	ID              uint64    `gorm:"primaryKey" json:"-"`
	HashCode        string    `gorm:"index:idx_act_hascode,unique,not null,size:37" json:"id"`
	FullName        string    `gorm:"not null" json:"fullName" validate:"min=3,max=80,alpha,required"`
	Identification  string    `gorm:"index:idx_identification,unique,not null" json:"identification" validate:"required"`
	Birth           string    `gorm:"not null" json:"birth" validate:"required"`
	AddressPostCode string    `gorm:"not null" json:"postCode" validate:"required"`
	Address         string    `gorm:"not null" json:"address" validate:"required"`
	AddressNumber   string    `gorm:"not null" json:"number" validate:"required"`
	AddressDetails  string    `gorm:"not null" json:"complement"`
	AddressNeighbor string    `gorm:"not null" json:"neighbor" validate:"required"`
	AddressCity     string    `gorm:"not null" json:"city" validate:"required"`
	AddressState    string    `gorm:"not null" json:"state" validate:"required"`
	AddressCountry  string    `gorm:"not null" json:"country" validate:"required"`
	PhoneNumber     string    `gorm:"not null" json:"phone" validate:"required"`
	Email           string    `gorm:"index:idx_act_email,unique,not null" json:"email" validate:"required,email"`
	Password        string    `gorm:"not null" json:"-" validate:"required"`
	CreatedAt       time.Time `gorm:"not null" json:"create_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"update_at"`
}

func (a Account) IsValid(checkAccountAlreadyExist func(email string, identification string) bool) (violations []string) {

	validate := validator.New()

	if error := validate.Struct(a); error != nil {
		for _, msg := range error.(validator.ValidationErrors) {
			violations = append(violations, msg.Error())
		}
	}

	if exist := checkAccountAlreadyExist(a.Email, a.Identification); exist {
		violations = append(violations, "Key: 'Account.Email or Account.Identification' Error: Any fields already exist check each.")
	}

	return violations
}
