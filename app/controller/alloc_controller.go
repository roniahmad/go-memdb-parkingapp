package controller

import (
	"github.com/roniahmad/parking-app/app/model"
	"github.com/roniahmad/parking-app/config"
)

type AllocController struct {
	Usecase model.LocAllocUsecase
	Config  *config.Config
}

func (c *AllocController) Park(number string) error {
	err := c.Usecase.Park(number)
	if err != nil {
		return err
	}

	return nil
}

func (c *AllocController) Leave(number string, hours int) error {
	err := c.Usecase.Leave(number, hours)
	if err != nil {
		return err
	}

	return nil
}

func (c *AllocController) Status() error {
	err := c.Usecase.Status()

	if err != nil {
		return err
	}

	return nil
}
