package repository

import (
	"github.com/hashicorp/go-memdb"
	"github.com/roniahmad/parking-app/app/model"
)

type lotRepository struct {
	Db *memdb.MemDB
}

// Create Parking Lot
func (l *lotRepository) CreateLot(number int) error {
	txn := l.Db.Txn(true)
	lot := &model.Lot{
		Number: number,
	}
	if err := txn.Insert("lot", lot); err != nil {
		return err
	}

	// Create lot allocation
	for i := range number {
		alloc := &model.LotAllocation{
			Lot:    i + 1,
			Number: " ",
			Color:  "",
			In:     0,
			Out:    0,
			Status: "empty",
		}

		if err := txn.Insert("lot_allocation", alloc); err != nil {
			return err
		}
	}
	txn.Commit()

	return nil
}

func NewLotRepository(db *memdb.MemDB) model.LotRepository {
	return &lotRepository{
		Db: db,
	}
}
