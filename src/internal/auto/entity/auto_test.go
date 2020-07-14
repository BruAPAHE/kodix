package entity_test

import (
	"testing"

	"kodix/src/internal/auto/entity"

	"github.com/stretchr/testify/assert"
)

func TestAuto_SetAttributes(t *testing.T) {
	failUseCase := struct {
		price  int
		status string
	}{
		-11,
		"test",
	}
	wellUseCase := struct {
		price  int
		status string
	}{
		11,
		"sold",
	}

	mockAutoEntity := entity.NewAuto("something", "something", 666)

	assert.Error(t, mockAutoEntity.SetStatus(failUseCase.status))
	assert.Error(t, mockAutoEntity.SetPrice(failUseCase.price))
	assert.NoError(t, mockAutoEntity.SetPrice(wellUseCase.price))
	assert.NoError(t, mockAutoEntity.SetStatus(wellUseCase.status))
}
