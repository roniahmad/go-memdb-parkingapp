package controller

import (
	"github.com/roniahmad/parking-app/app/model"
	"github.com/roniahmad/parking-app/config"
)

type LotController struct {
	Usecase model.LotUsecase
	Config  *config.Config
}

func (c *LotController) CreateParkingLot(number int) error {
	err := c.Usecase.CreateParkingLot(number)
	if err != nil {
		return err
	}

	return nil
}
