package repository

import (
	"github.com/hashicorp/go-memdb"
	"github.com/roniahmad/parking-app/app/model"
)

type allocationRepository struct {
	Db *memdb.MemDB
}

// Find Empty Slot
func (l *allocationRepository) FindNearestEmptySlot() (interface{}, error) {
	txn := l.Db.Txn(false)
	defer txn.Abort()

	// find empty first slot
	raw, err := txn.First("lot_allocation", "status", "empty")
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, err
	}

	return raw, err
}

// GetAll implements model.LotAllocRepository.
func (l *allocationRepository) GetAll() (memdb.ResultIterator, error) {
	txn := l.Db.Txn(false)
	defer txn.Abort()

	lots, err := txn.Get("lot_allocation", "id")
	if err != nil {
		return nil, err
	}

	return lots, nil
}

// Find
func (l *allocationRepository) IsExist(carNumber string) (interface{}, error) {
	txn := l.Db.Txn(false)
	defer txn.Abort()

	// query by car number
	raw, err := txn.First("lot_allocation", "number", carNumber)
	if err != nil {
		return nil, err
	}

	return raw, err
}

// Create
func (l *allocationRepository) Create(alloc *model.LotAllocation) error {
	txn := l.Db.Txn(true)
	if err := txn.Insert("lot_allocation", alloc); err != nil {
		return err
	}
	txn.Commit()

	return nil
}

// Delete
func (l *allocationRepository) Delete(alloc *model.LotAllocation) error {
	txn := l.Db.Txn(true)
	if err := txn.Delete("lot_allocation", alloc); err != nil {
		return err
	}
	txn.Commit()

	return nil
}

// Update
func (l *allocationRepository) Update(alloc *model.LotAllocation) error {
	txn := l.Db.Txn(true)
	// since modifiying value-inplace in memdb is not supported,
	// we make a copy and delete exisiting then make a new insert
	newAlloc := &model.LotAllocation{
		Lot:    alloc.Lot,
		Number: alloc.Number,
		Color:  alloc.Color,
		In:     alloc.In,
		Out:    alloc.Out,
		Status: alloc.Status,
	}

	if err := txn.Delete("lot_allocation", alloc); err != nil {
		return err
	}

	if err := txn.Insert("lot_allocation", newAlloc); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func NewAllocationRepository(db *memdb.MemDB) model.LotAllocRepository {
	return &allocationRepository{
		Db: db,
	}
}
