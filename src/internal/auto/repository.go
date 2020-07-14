package auto

import (
	"context"

	"kodix/src/internal/models"
)

type IAutoRepository interface {
	GetById(ctx context.Context, id int) (*models.AutoAttributes, error)
	GetAll(ctx context.Context) ([]*models.AutoAttributes, error)
	Create(ctx context.Context, auto models.AutoAttributes) error
	Update(ctx context.Context, auto models.AutoAttributes, id int) error
	Delete(ctx context.Context, id int) error
}
