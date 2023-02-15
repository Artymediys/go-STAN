package handlers

import (
	"go_STAN/internal/db"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var dataBase *db.DataBase

func ResultPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/web/templates/result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var id int = 1
	if r.Method == "POST" {
		idSTR := r.FormValue("id")
		id, err = strconv.Atoi(idSTR)
		if err != nil {
			log.Println(err)
			return
		}
	}

	var result *db.MainInfo = searchInfo(id)

	data := struct {
		Result *db.MainInfo
	}{
		Result: result,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func searchInfo(id int) *db.MainInfo {
	if info, ok := dataBase.Cache.Get(id); ok {
		return &info
	}
	return nil
}

func GetDBObject(db *db.DataBase) {
	dataBase = db
}
