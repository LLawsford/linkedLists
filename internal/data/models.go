package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	ItemList ItemListModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		ItemList: ItemListModel{DB: db},
	}
}
