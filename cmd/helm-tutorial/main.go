package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/Hrtkk/helm-tutorial/pkg/api"
	"github.com/gorilla/mux"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", makeHandler(api.HomeHandler))
	r.HandleFunc("/articles", makeHandler(api.ArticlesHandler))
	r.HandleFunc("/products", makeHandler(api.ProductsHandler))
	r.HandleFunc("/view/", makeHandler(api.ViewHandler))
	r.HandleFunc("/edit", makeHandler(api.EditHandler))
	r.HandleFunc("/save", makeHandler(api.SaveHandler))

	log.Fatal(http.ListenAndServe(":8085", r))
}
