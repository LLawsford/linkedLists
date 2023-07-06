package itemLists

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AttachItemListsRoutes(r *chi.Mux) *chi.Mux {
	(*r).Route("/api/v1/itemLists", func(r chi.Router) {
		r.Get("/{itemListId}", func(writer http.ResponseWriter, request *http.Request) {
			itemListId, err := strconv.ParseInt(chi.URLParam(request, "itemListId"), 10, 32)
			if err != nil {
				http.Error(writer, "cannot parse id", http.StatusBadRequest)
			}

			fmt.Fprintf(writer, "getting item list by id: %d", itemListId)
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("getting item lists"))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("creating item lists"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("deleting item list"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("deleting item list"))
		})
	})

	return r
}
