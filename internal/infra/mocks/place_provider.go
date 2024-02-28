package mocks

import (
	"github.com/leonardfreitas/golang-cloud-run/internal/entity"
	"github.com/leonardfreitas/golang-cloud-run/internal/infra/place"
	"github.com/stretchr/testify/mock"
)

var _ place.PlaceProviderInterface = (*PlaceProviderMock)(nil)

type PlaceProviderMock struct {
	mock.Mock
}

func NewPlaceProviderMock() *PlaceProviderMock {
	return &PlaceProviderMock{}
}

func (lm *PlaceProviderMock) GetByCep(cep string) (entity.Place, error) {
	args := lm.Called(cep)
	return args.Get(0).(entity.Place), args.Error(1)
}
