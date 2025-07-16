package internal

import (
	"tesodev-korpes/CustomerService/internal/types"
)

// here is an example what this helper method that does data casting from db model to response model
// the return statement that I commented out repreents an introduction that how you can implement it
// you can delete after you'd completed the helper method, its a placeholder put here just to prevent getting errors at
// the beginning
func ToCustomerResponse(customer *types.Customer) *types.CustomerResponseModel {
	return &types.CustomerResponseModel{
		Id:          customer.Id,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Addresses:   customer.Addresses,
		DateOfBirth: customer.DateOfBirth,
		Gender:      customer.Gender,
		Notes:       customer.Notes,
	}
}
