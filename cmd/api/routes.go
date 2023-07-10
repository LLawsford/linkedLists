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
		r.Post("/", app.createItemListHandler)
		r.Get("/", app.dummyHandler)
		r.Delete("/", app.dummyHandler)
		r.Put("/", app.dummyHandler)
		r.Patch("/", app.dummyHandler)
	})

	// v1 items routes
	router.Route("/api/v1/itemsLists/{itemListId}/items", func(r chi.Router) {
		r.Get("/{itemId}", app.dummyHandler)
		r.Get("/", app.dummyHandler)
		r.Post("/", app.dummyHandler)
		r.Delete("/", app.dummyHandler)
		r.Put("/", app.dummyHandler)
		r.Patch("/", app.dummyHandler)
	})
	return router
}
