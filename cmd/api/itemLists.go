package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/LLawsford/linkedLists/internal/data"
	"github.com/LLawsford/linkedLists/internal/validator"
	"github.com/go-chi/chi"
)

func (app *application) createItemListHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	itemList := &data.ItemList{
		Title:       input.Title,
		Description: input.Description,
	}

	v := validator.New()

	if data.ValidateItemList(v, itemList); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showItemListHandler(w http.ResponseWriter, r *http.Request) {
	itemListId, err := strconv.ParseInt(chi.URLParam(r, "itemListId"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	itemList := data.ItemList{
		ID:          itemListId,
		CreatedAt:   time.Now(),
		Title:       "some random title",
		Description: "some random description",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"itemList": itemList}, nil)
	if err != nil {
		http.Error(w, "The server encountered problem", http.StatusInternalServerError)
	}
}
