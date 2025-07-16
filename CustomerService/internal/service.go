package internal

import (
	"context"
	"tesodev-korpes/CustomerService/internal/types"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByID(ctx context.Context, id string) (*types.Customer, error) {
	customer, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) Create(ctx context.Context, customer *types.CustomerRequestModel) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dbCustomer := &types.Customer{
		FirstName:    customer.FirstName,
		LastName:     customer.LastName,
		Email:        customer.Email,
		Phone:        customer.Phone,
		Addresses:    customer.Addresses,
		DateOfBirth:  customer.DateOfBirth,
		Gender:       customer.Gender,
		Notes:        customer.Notes,
		Preferences:  customer.Preferences,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.repo.Create(ctx, dbCustomer)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, id string, updatedCustomer types.CustomerRequestModel) error {
	return s.repo.Update(ctx, id, updatedCustomer)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
