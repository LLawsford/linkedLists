package items

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AttachItemsRoutes(r *chi.Mux) *chi.Mux {
	(*r).Route("/api/v1/itemsLists/{itemListId}/items", func(r chi.Router) {
		r.Get("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("getting item by id on item list"))
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("getting items on item list"))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("creating items on item list"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("deleting items on item list"))
		})
	})

	return r
}
