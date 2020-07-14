package mock

import (
	"context"

	"kodix/src/internal/auto/entity"
	"kodix/src/internal/models"

	"github.com/stretchr/testify/mock"
)

type AutoRepositoryMock struct {
	mock.Mock
}

func (a AutoRepositoryMock) GetById(ctx context.Context, id int) (*models.AutoAttributes, error) {
	panic("implement me")
}

func (a AutoRepositoryMock) GetAll(ctx context.Context) ([]*models.AutoAttributes, error) {
	panic("implement me")
}

func (a AutoRepositoryMock) Create(ctx context.Context, auto entity.Auto) error {
	panic("implement me")
}

func (a AutoRepositoryMock) Update(ctx context.Context, auto entity.Auto, id int) error {
	panic("implement me")
}

func (a AutoRepositoryMock) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

