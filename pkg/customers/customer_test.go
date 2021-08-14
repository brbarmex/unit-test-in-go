package customers_test

import (
	"testing"

	"github.com/BrunoBMelo/shoestore/pkg/customers"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateStateStruct_Customer(t *testing.T) {
	cases := []struct {
		entity         customers.Customer
		expect         []string
		lengthExpected int
	}{
		{customers.Customer{}, nil, 13},
		{customers.Customer{FullName: "Bruno"}, nil, 12},
		{customers.Customer{FullName: "Bruno", Identification: "a"}, nil, 11},
		{customers.Customer{FullName: "Bruno", Identification: "a", Birth: "30071994"}, nil, 10},
	}

	for _, c := range cases {
		got := len(c.entity.IsValid())
		assert.Equal(t, c.lengthExpected, got)
	}
}

// func Test_Create_Customer(t *testing.T) {

// 	cases := []struct {
// 		obj          customers.Customer
// 		existingMail bool
// 		errExpected  error
// 		msg          string
// 	}{
// 		{customers.Customer{}, true, errors.New("there is an account using this email"), "there is an account using this email"},
// 		{customers.Customer{}, false, nil, ""},
// 	}

// 	for _, c := range cases {
// 		mock := &mocks.CustomerRepository{}

// 		mock.
// 			On("CheckIfMailAlredyExist", c.obj.Email).Return(c.existingMail).
// 			On("Add", &c.obj).Return(c.errExpected).Maybe()

// 		service := customers.CustomerServices{Repository: mock}

// 		_, err := service.CreateCustomer(&c.obj)

// 		assert.EqualError(t, err, c.msg)

// 		mock.AssertExpectations(t)
// 	}
// }
