package bootstrap

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

/*
	type Lot struct {
		Number int
	}

	type LotAllocation struct {
		Lot    int
		Number string
		Color  string
		In     int
		Out    int
		Status string
	}
*/
func initializeSchema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"lot": &memdb.TableSchema{
				Name: "lot",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Number"},
					},
				},
			},
			"lot_allocation": &memdb.TableSchema{
				Name: "lot_allocation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Lot"},
					},
					"number": &memdb.IndexSchema{
						Name:    "number",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Number"},
					},
					// "color": &memdb.IndexSchema{
					// 	Name:    "color",
					// 	Unique:  false,
					// 	Indexer: &memdb.StringFieldIndex{Field: "Color"},
					// },
					"status": &memdb.IndexSchema{
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
				},
			},
		},
	}

	return schema
}

func NewDb() *memdb.MemDB {
	schema := initializeSchema()

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}
