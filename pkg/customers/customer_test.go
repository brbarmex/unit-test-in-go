package customers_test

import (
	"testing"

	"github.com/BrunoBMelo/shoestore/pkg/customers/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_validate_state_of_struct_customer(t *testing.T) {
	for _, usecase := range mocks.Cases_possibles_states_of_structs {
		got := usecase.Entity.IsValid()
		assert.Equal(t, len(usecase.ResultExpected), len(got))
		assert.ElementsMatch(t, usecase.ResultExpected, got)
	}
}