package weather

import "github.com/leonardfreitas/golang-cloud-run/internal/entity"

type WeatherProviderInterface interface {
	GetWeather(city string) (entity.Weather, error)
}
