package usecase

import (
	"fmt"

	"github.com/roniahmad/parking-app/app/model"
	"github.com/roniahmad/parking-app/config"
	"github.com/roniahmad/parking-app/internal/helper"
)

type lotAllocationUsecase struct {
	repo   model.LotAllocRepository
	config *config.Config
}

// Status implements model.LocAllocUsecase.
func (l *lotAllocationUsecase) Status() error {
	lots, err := l.repo.GetAll()
	if err != nil {
		return err
	}

	fmt.Println("Slot No. Registration No.")
	for lot := lots.Next(); lot != nil; lot = lots.Next() {
		p := lot.(*model.LotAllocation)
		fmt.Printf("%d  %s\n", p.Lot, p.Number)
	}

	return nil
}

// Car Leave
func (l *lotAllocationUsecase) Leave(carNumber string, hours int) error {
	// check if car is registered before
	raw, err := l.repo.IsExist(carNumber)
	if err != nil {
		fmt.Printf("Registration number %s not found\n", carNumber)
		return fmt.Errorf("registration number %s not found", carNumber)
	}

	if raw == nil {
		fmt.Printf("Registration number %s not found\n", carNumber)
		return fmt.Errorf("registration number %s not found", carNumber)
	}

	// update car allocation
	if err = l.repo.Update(&model.LotAllocation{
		Lot:    raw.(*model.LotAllocation).Lot,
		Number: " ",
		Color:  "",
		In:     0,
		Out:    0,
		Status: "empty",
	}); err != nil {
		return err
	}

	// calculate bills
	var (
		bill               = 0
		max                = l.config.MaxFirstHour
		chargeFirstTwoHour = l.config.ChargeFirstTwoHours
		chargeNextHour     = l.config.ChargeNextHours
	)

	if hours <= max {
		bill = chargeFirstTwoHour
	} else {
		bill = chargeFirstTwoHour + ((hours - max) * chargeNextHour)
	}
	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d \n", carNumber, raw.(*model.LotAllocation).Lot, bill)

	return nil
}

// Park the car
func (l *lotAllocationUsecase) Park(carNumber string) error {
	// check if car is registered before
	raw, _ := l.repo.IsExist(carNumber)

	if raw != nil {
		fmt.Printf("Car number %s already registered \n", carNumber)
		return fmt.Errorf("car number %s already registered", carNumber)
	}

	// find empty slot
	raw, err := l.repo.FindNearestEmptySlot()
	if err != nil {
		return err
	}

	if raw == nil {
		fmt.Println("Sorry, parking lot is full")
		return fmt.Errorf("sorry, parking lot is full")
	}

	// if there is empty slot
	if err = l.repo.Create(&model.LotAllocation{
		Lot:    raw.(*model.LotAllocation).Lot,
		Number: carNumber,
		Color:  helper.GenerateRandomCardColor(),
		In:     0,
		Out:    0,
		Status: "filled",
	}); err != nil {

		return err
	}
	fmt.Printf("Allocated Slot Number: %d\n", raw.(*model.LotAllocation).Lot)
	return nil
}

func NewLotAllocationUsecase(repo model.LotAllocRepository, config *config.Config) model.LocAllocUsecase {
	return &lotAllocationUsecase{
		repo:   repo,
		config: config,
	}
}
