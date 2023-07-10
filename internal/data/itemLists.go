package data

import (
	"time"

	"github.com/LLawsford/linkedLists/internal/validator"
)

type ItemList struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func ValidateItemList(v *validator.Validator, itemList *ItemList) {
	v.Check(itemList.Title != "", "title", "must be provided")
	v.Check(len(itemList.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(itemList.Description != "", "description", "must be provided")
	v.Check(len(itemList.Description) <= 1500, "description", "must not be more than 1500 bytes long")
}
