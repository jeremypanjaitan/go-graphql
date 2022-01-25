package app

import (
	"errors"
	"go-graphql/delivery"
	"go-graphql/infra"
	"go-graphql/manager"
	"go-graphql/repo"
	"go-graphql/usecase"
	"os"

	"github.com/joho/godotenv"
)

type AppEntity interface {
	Run()
}

type App struct {
	httpDelivery delivery.HttpDeliveryEntity
}

func init() {

	//Check if .env file is exist or not
	if _, err := os.Stat(".env"); !errors.Is(err, os.ErrNotExist) {
		godotenv.Load()
	}
}

func NewApp() AppEntity {
	app := new(App)

	httpInfra := infra.NewHttpInfra()
	productRepo := repo.NewProductRepo()
	gqlManager := manager.NewGraphqlManager(productRepo)
	graphqlUseCase := usecase.NewGraphqlUsecase(gqlManager.GetGraphql())
	httpDelivery := delivery.NewHttpDelivery(httpInfra, graphqlUseCase)
	app.httpDelivery = httpDelivery

	return app
}

func (a *App) Run() {
	a.httpDelivery.RunHttpDelivery()
}
