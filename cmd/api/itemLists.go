// TODO: currently patch and put request here has some overlapping logic - could be limited to only one request type
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	err = app.models.ItemList.Insert(itemList)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/api/v1/itemLists/%d", itemList.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"itemList": itemList}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showItemListHandler(w http.ResponseWriter, r *http.Request) {
	itemListId, err := strconv.ParseInt(chi.URLParam(r, "itemListId"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	itemList, err := app.models.ItemList.Get(itemListId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"itemList": itemList}, nil)
	if err != nil {
		http.Error(w, "The server encountered problem", http.StatusInternalServerError)
	}
}

func (app *application) updateItemListHandler(w http.ResponseWriter, r *http.Request) {
	itemListId, err := strconv.ParseInt(chi.URLParam(r, "itemListId"), 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	itemList, err := app.models.ItemList.Get(itemListId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	log.Printf("%+v\n", itemList)

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	itemList.Title = input.Title
	itemList.Description = input.Description

	v := validator.New()

	if data.ValidateItemList(v, itemList); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.ItemList.Update(itemList)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(itemList.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"itemList": itemList}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteItemListHandler(w http.ResponseWriter, r *http.Request) {
	itemListId, err := strconv.ParseInt(chi.URLParam(r, "itemListId"), 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.ItemList.Delete(itemListId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "item list successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) partiallyUpdateItemListHandler(w http.ResponseWriter, r *http.Request) {
	itemListId, err := strconv.ParseInt(chi.URLParam(r, "itemListId"), 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	itemList, err := app.models.ItemList.Get(itemListId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		itemList.Title = *input.Title
	}

	if input.Description != nil {
		itemList.Description = *input.Description
	}

	v := validator.New()

	if data.ValidateItemList(v, itemList); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.ItemList.Update(itemList)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(itemList.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"itemList": itemList}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
