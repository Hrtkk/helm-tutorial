package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Hrtkk/helm-tutorial/models"
	"github.com/Hrtkk/helm-tutorial/pkg/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request, title string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Home handler")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request, title string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Article handler")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request, title string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Article handler")
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	utils.RenderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	utils.RenderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
