package auto

import (
	"context"

	"kodix/src/internal/auto/entity"
	"kodix/src/internal/models"
)

type IAutoService interface {
	GetAutoById(ctx context.Context, id int) (*models.AutoAttributes, error)
	GetAllAuto(ctx context.Context) ([]*models.AutoAttributes, error)
	CreateAuto(ctx context.Context, auto entity.Auto) error
	UpdateAuto(ctx context.Context, auto entity.Auto, id int) error
	DeleteAuto(ctx context.Context, id int) error
}
