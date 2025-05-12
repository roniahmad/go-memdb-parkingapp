package model

type Lot struct {
	Number int
}

type LotRepository interface {
	CreateLot(number int) error
}

type LotUsecase interface {
	CreateParkingLot(number int) error
}
