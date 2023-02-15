package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
