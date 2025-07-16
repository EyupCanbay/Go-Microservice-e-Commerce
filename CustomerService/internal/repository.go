package internal

import (
	"context"
	"errors"
	"tesodev-korpes/CustomerService/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{
		collection: col,
	}
}

func (r *Repository) FindByID(ctx context.Context, id string) (*types.Customer, error) {
	var customer *types.Customer
	return customer, nil
}

func (r *Repository) Create(ctx context.Context, customer *types.Customer) error {
	// Placeholder method
	return nil
}

func (r *Repository) Update(ctx context.Context, id string, update types.CustomerRequestModel) error {
	// Placeholder method
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("customer not found")
	}

	return nil
}

func (r *Repository) GetList(ctx context.Context) ([]types.Customer, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customers []types.Customer
	for cursor.Next(ctx) {
		var customer types.Customer
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
