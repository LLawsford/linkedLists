package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	// TODO: when needed, can be replaced with custom logger
	router.Use(middleware.Logger)

	// errors
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	// v1 healthcheck
	router.Get("/api/v1/healthcheck", app.healthcheckHandler)

	// v1 item lists routes
	router.Route("/api/v1/itemLists", func(r chi.Router) {
		r.Get("/{itemListId}", app.showItemListHandler)
		r.Put("/{itemListId}", app.updateItemListHandler)
		r.Delete("/{itemListId}", app.deleteItemListHandler)
		r.Patch("/{itemListId}", app.partiallyUpdateItemListHandler)
		r.Post("/", app.createItemListHandler)
		r.Get("/", app.dummyHandler)
	})

	return router
}
