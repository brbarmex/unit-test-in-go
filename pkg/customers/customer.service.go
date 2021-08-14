package customers

import "errors"

func (s *CustomerServices) CreateCustomer(customer *Customer) ([]string, error) {

	inconsistencies_services := []string{}

	if inconsistencies := customer.IsValid(); len(inconsistencies) > 0 {
		inconsistencies_services = append(inconsistencies_services, inconsistencies...)
	}

	emailExists, err := s.Repository.CheckIfMailAlredyExist(customer.Email)
	if err != nil {
		return inconsistencies_services, errors.New("an error occurred while verifying the e-mail already exists")
	}

	if emailExists {
		inconsistencies_services = append(inconsistencies_services, "the e-mail already exist in other account")
	}

	identificationExists, err := s.Repository.CheckIfIdentificationExist(customer.Identification)

	if err != nil {
		return inconsistencies_services, errors.New("an error occurred while verifying the identification already exists")
	}

	if identificationExists {
		inconsistencies_services = append(inconsistencies_services, "the identification already exist in other account")
	}

	if err := s.Repository.Save(customer); err != nil {
		return inconsistencies_services, errors.New("an error occurred while save customer in database")
	}

	return inconsistencies_services, nil
}
