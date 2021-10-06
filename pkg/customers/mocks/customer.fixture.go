package mocks

import customers "github.com/BrunoBMelo/shoestore/pkg/customers"

var Cases_possibles_states_of_structs = []struct {
	Entity         customers.Customer
	ResultExpected []string
}{
	{customers.Customer{
		Id:              0,
		UUID:            "",
		FullName:        "Bruno",
		Identification:  "123456789",
		Birth:           "10102001",
		AddressPostCode: "09930442",
		Address:         "STREET A",
		AddressNumber:   "21A",
		AddressDetails:  "",
		AddressNeighbor: "BROWN",
		AddressCity:     "SAO PAULO",
		AddressState:    "SP",
		AddressCountry:  "BRAZIL",
		PhoneNumber:     "11995283697",
		Email:           "BRUNO",
		Password:        "USAHAUS",
		CreatedAt:       "",
		UpdatedAt:       "",
	}, []string{}},
}

// cases := []struct {
// 	entity         customers.Customer
// 	expect         []string
// }{
// 	{customers.Customer{}, nil, 13},
// 	{customers.Customer{FullName: "Bruno"}, nil, 12},
// 	{customers.Customer{FullName: "Bruno", Identification: "a"}, nil, 11},
// 	{customers.Customer{FullName: "Bruno", Identification: "a", Birth: "30071994"}, nil, 10},
// }
