package usecase

import (
	"github.com/roniahmad/parking-app/app/model"
	"github.com/roniahmad/parking-app/config"
)

type lotUsecase struct {
	repo   model.LotRepository
	config *config.Config
}

// Create Parking Lot
func (l *lotUsecase) CreateParkingLot(number int) error {
	err := l.repo.CreateLot(number)
	return err
}

func NewLotUsecase(repo model.LotRepository, config *config.Config) model.LotUsecase {
	return &lotUsecase{
		repo:   repo,
		config: config,
	}
}
