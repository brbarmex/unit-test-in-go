package customers

import "github.com/go-playground/validator/v10"

type Customer struct {
	Id              uint64 `json:"-"`
	UUID            string `json:"id"`
	FullName        string `json:"fullName" validate:"required"`
	Identification  string `json:"identification" validate:"required"`
	Birth           string `json:"birth" validate:"required"`
	AddressPostCode string `json:"zipCode" validate:"required"`
	Address         string `json:"street" validate:"required"`
	AddressNumber   string `json:"streeNumber" validate:"required"`
	AddressDetails  string `json:"streetDetail" `
	AddressNeighbor string `json:"neighbor" validate:"required"`
	AddressCity     string `json:"city" validate:"required"`
	AddressState    string `json:"state" validate:"required"`
	AddressCountry  string `json:"country" validate:"required"`
	PhoneNumber     string `json:"phone" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"-" validate:"required"`
	CreatedAt       string `json:"create_at"`
	UpdatedAt       string `json:"update_at"`
}

func (c *Customer) IsValid() []string {
	validate := validator.New()
	violations := []string{}
	if error := validate.Struct(c); error != nil {
		for _, msg := range error.(validator.ValidationErrors) {
			violations = append(violations, msg.Error())
		}
	}
	return violations
}

type CustomerServices struct {
	Repository CustomerRepository
}

type CustomerRepository interface {
	Save(customer *Customer) error
	CheckIfMailAlredyExist(mail string) (bool, error)
	CheckIfIdentificationExist(identification string) (bool, error)
}
