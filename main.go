// TODO: json encoding + decoding
// TODO: json validators
// TODO: database connection
// TODO: dockerization
// TODO: JWT basic validation
// TODO: tests
// TODO: ci/cd
// TODO: SSL
package main

import (
	"net/http"

	"github.com/LLawsford/linkedLists/itemLists"
	"github.com/LLawsford/linkedLists/items"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	port string
}

func (apiServer *APIServer) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	itemLists.AttachItemListsRoutes(r)
	items.AttachItemsRoutes(r)

	//TODO research ListenAndServeTLS for https
	http.ListenAndServe(apiServer.port, r)
}

func main() {
	apiServer := APIServer{
		//TODO: get variables from env file
		port: ":3030",
	}

	apiServer.Run()
}
