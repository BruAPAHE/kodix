package services

import (
	"context"
	"kodix/src/internal/auto"

	"kodix/src/internal/auto/entity"
	"kodix/src/internal/models"
)

type AutoService struct {
	repo auto.IAutoRepository
}

func NewAutoService(repo auto.IAutoRepository) *AutoService {
	return &AutoService{repo: repo}
}

func (a AutoService) GetAutoById(ctx context.Context, id int) (*models.AutoAttributes, error) {
	return a.repo.GetById(ctx, id)
}

func (a AutoService) GetAllAuto(ctx context.Context) ([]*models.AutoAttributes, error) {
	return a.repo.GetAll(ctx)
}

func (a AutoService) CreateAuto(ctx context.Context, auto entity.Auto) error {
	return a.repo.Create(ctx, models.AutoAttributes{
		Brand:   auto.Brand(),
		Model:   auto.Model(),
		Price:   uint(auto.Price()),
		Status:  auto.Status(),
		Mileage: auto.Mileage(),
	})
}

func (a AutoService) UpdateAuto(ctx context.Context, auto entity.Auto, id int) error {
	return a.repo.Update(ctx, models.AutoAttributes{
		Brand:   auto.Brand(),
		Model:   auto.Model(),
		Price:   uint(auto.Price()),
		Status:  auto.Status(),
		Mileage: auto.Mileage(),
	}, id)
}

func (a AutoService) DeleteAuto(ctx context.Context, id int) error {
	return a.repo.Delete(ctx, id)
}
