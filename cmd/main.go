package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardfreitas/golang-cloud-run/configs"
	"github.com/leonardfreitas/golang-cloud-run/internal/infra/place/viacep"
	"github.com/leonardfreitas/golang-cloud-run/internal/infra/weather/weatherapi"
	"github.com/leonardfreitas/golang-cloud-run/internal/usecase"
	"github.com/leonardfreitas/golang-cloud-run/internal/web/handler"
)

func getConfig() *configs.Config {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Erro ao obter informações do arquivo.")
	}
	goDir := filepath.Dir(currentFile)

	config, err := configs.LoadConfig(goDir)
	if err != nil {
		panic(err)
	}
	return config
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	viaCep := viacep.NewViaCep()
	config := getConfig()
	weatherApi := weatherapi.NewWeatherAPI(config)
	placeUseCase := usecase.NewGetPlaceForecastUseCase(viaCep, weatherApi)

	getPlaceTemperaturesHandler := handler.NewGetPlaceTemperaturesHandler(placeUseCase)
	router.MethodFunc(http.MethodGet, "/", getPlaceTemperaturesHandler.Handle)

	log.Println("Iniciando o servidor web...")
	http.ListenAndServe(":8080", router)
}
