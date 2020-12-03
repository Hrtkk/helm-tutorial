package main

import (
	"log"
	"net/http"

	"github.com/Hrtkk/helm-tutorial/pkg/api"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", api.HomeHandler)
	r.HandleFunc("/articles", api.ArticlesHandler)
	r.HandleFunc("/products", api.ProductsHandler)
	log.Fatal(http.ListenAndServe(":8085", r))
}
