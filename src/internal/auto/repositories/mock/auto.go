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
	args := a.Called(ctx, id)

	return args.Get(0).(*models.AutoAttributes), args.Error(1)
}

func (a AutoRepositoryMock) GetAll(ctx context.Context) ([]*models.AutoAttributes, error) {
	args := a.Called(ctx)

	return args.Get(0).([]*models.AutoAttributes), args.Error(1)
}

func (a AutoRepositoryMock) Create(ctx context.Context, auto entity.Auto) error {
	args := a.Called(ctx, auto)

	return args.Error(1)
}

func (a AutoRepositoryMock) Update(ctx context.Context, auto entity.Auto, id int) error {
	args := a.Called(ctx, auto, id)

	return args.Error(1)
}

func (a AutoRepositoryMock) Delete(ctx context.Context, id int) error {
	args := a.Called(ctx, id)

	return args.Error(1)
}
