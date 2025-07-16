package internal

import (
	"context"
	"errors"
	"tesodev-korpes/CustomerService/internal/types"
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

	//challenge (everything should be observable somehow in the response or console (print)):
	// 1) do something with using for loop by using customer model and manipulate it (you can add an additional field for it)
	// 2) do something with switch-case
	// 3) do something with goroutines (you should give us an example for both scenarios of not using goroutines and using)
	// 3.1) calculate the elapsed time for both scenarios and show us the gained time
	// 4) add an additional field and use maps
	// 5) add an additional field and use arrays
	// 6) manipulate an existing data to see how pointers and values work
	return customer, nil
}

func (s *Service) Create(ctx context.Context, customer interface{}) error {
	return s.repo.Create(ctx, customer)
}

func (s *Service) Update(ctx context.Context, id string, update interface{}) error {
	updateMap, ok := update.(map[string]interface{})
	if !ok {
		return errors.New("invalid update payload")
	}
	return s.repo.Update(ctx, id, updateMap)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
