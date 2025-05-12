package bootstrap

import (
	"github.com/hashicorp/go-memdb"
	"github.com/roniahmad/parking-app/app/controller"
	"github.com/roniahmad/parking-app/app/repository"
	"github.com/roniahmad/parking-app/app/usecase"
	"github.com/roniahmad/parking-app/config"
)

type Application struct {
	Config          *config.Config
	Db              *memdb.MemDB
	LotController   *controller.LotController
	AllocController *controller.AllocController
}

func NewApp() *Application {
	app := &Application{}
	app.Config = NewConfig()
	app.Db = NewDb()

	lotRepository := repository.NewLotRepository(app.Db)
	allocRepository := repository.NewAllocationRepository(app.Db)

	ucLot := usecase.NewLotUsecase(lotRepository, app.Config)
	ucAlloc := usecase.NewLotAllocationUsecase(allocRepository, app.Config)

	app.LotController = &controller.LotController{
		Usecase: ucLot,
		Config:  app.Config,
	}

	app.AllocController = &controller.AllocController{
		Usecase: ucAlloc,
		Config:  app.Config,
	}

	return app
}
