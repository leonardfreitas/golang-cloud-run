package place

import "github.com/leonardfreitas/golang-cloud-run/internal/entity"

type PlaceProviderInterface interface {
	GetByCep(cep string) (entity.Place, error)
}
