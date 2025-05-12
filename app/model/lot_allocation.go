package model

import (
	"github.com/hashicorp/go-memdb"
)

type LotAllocation struct {
	Lot    int
	Number string
	Color  string
	In     int
	Out    int
	Status string
}

type LotAllocRepository interface {
	Create(alloc *LotAllocation) error
	Update(alloc *LotAllocation) error
	Delete(alloc *LotAllocation) error
	IsExist(carNumber string) (interface{}, error)
	GetAll() (memdb.ResultIterator, error)
	FindNearestEmptySlot() (interface{}, error)
}

type LocAllocUsecase interface {
	Park(number string) error
	Leave(carNumber string, hours int) error
	Status() error
}
