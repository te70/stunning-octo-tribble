package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/te70/stunning-octo-tribble/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.handle("/", r)
	log.Fatal(http.ListenAndServe("localhost: 8000", r))
}
