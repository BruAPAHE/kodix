package repositories

import (
	"context"
	"fmt"
	"os"

	"kodix/src/internal/auto/entity"
	"kodix/src/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutoRepository struct {
	pool *mongo.Collection
}

var collection = os.Getenv("DB_COLLECTION")

func NewAutoRepository(connect *mongo.Database) *AutoRepository {
	return &AutoRepository{
		pool: connect.Collection(collection),
	}
}

func (a AutoRepository) GetById(ctx context.Context, id int) (*models.AutoAttributes, error) {
	attributes := new(models.AutoAttributes)

	if err := a.pool.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(attributes); err != nil {
		return nil, fmt.Errorf("server error")
	}

	return attributes, nil
}

func (a AutoRepository) GetAll(ctx context.Context) ([]*models.AutoAttributes, error) {
	var result []*models.AutoAttributes

	rows, err := a.pool.Find(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("server error")
	}

	for rows.Next(ctx) {
		attributes := new(models.AutoAttributes)

		err := rows.Decode(attributes)
		if err != nil {
			return nil, err
		}
		result = append(result, attributes)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("server error")
	}
	return result, nil
}

func (a AutoRepository) Create(ctx context.Context, auto models.AutoAttributes) error {
	attributes := &models.AutoAttributes{
		Brand:   auto.Brand(),
		Model:   auto.Model(),
		Price:   uint(auto.Price()),
		Status:  auto.Status(),
		Mileage: auto.Mileage(),
	}
	_, err := a.pool.InsertOne(ctx, attributes)
	if err != nil {
		return fmt.Errorf("server error")
	}

	return nil
}

func (a AutoRepository) Update(ctx context.Context, auto models.AutoAttributes, id int) error {
	panic("implement me")
}

func (a AutoRepository) Delete(ctx context.Context, id int) error {
	_, err := a.pool.DeleteOne(ctx, bson.M{
		"_id": id,
	})
	return err
}
