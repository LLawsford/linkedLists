package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/LLawsford/linkedLists/internal/validator"
)

type ItemList struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Version     int32     `json:"version"`
}

func ValidateItemList(v *validator.Validator, itemList *ItemList) {
	v.Check(itemList.Title != "", "title", "must be provided")
	v.Check(len(itemList.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(itemList.Description != "", "description", "must be provided")
	v.Check(len(itemList.Description) <= 1500, "description", "must not be more than 1500 bytes long")
}

type ItemListModel struct {
	DB *sql.DB
}

func (il ItemListModel) Insert(itemList *ItemList) error {
	query := `
		INSERT INTO item_lists (title, description)
		VALUES ($1, $2)
		RETURNING id, created_at, version
	`
	args := []any{itemList.Title, itemList.Description}

	return il.DB.QueryRow(query, args...).Scan(&itemList.ID, &itemList.CreatedAt)
}

func (il ItemListModel) Get(id int64) (*ItemList, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, title, description
		FROM item_lists
		WHERE id = $1
	`

	var itemList ItemList

	err := il.DB.QueryRow(query, id).Scan(
		&itemList.ID,
		&itemList.CreatedAt,
		&itemList.Title,
		&itemList.Description,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &itemList, nil
}

func (il ItemListModel) Update(itemList *ItemList) error {
	query := `
		UPDATE item_lists
		SET title = $1, description = $2, version = version + 1
		WHERE id = $3 AND version = $6
		RETURNING version
	`

	args := []any{
		itemList.Title,
		itemList.Description,
		itemList.ID,
		itemList.Version,
	}

	err := il.DB.QueryRow(query, args...).Scan(&itemList.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (il ItemListModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM item_lists
		WHERE id = $1
	`
	result, err := il.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
