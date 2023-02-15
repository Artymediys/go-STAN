package web

import (
	"context"
	"go_STAN/internal/db"
	"go_STAN/internal/web/handlers"
	"log"
	"net/http"
)

type Web struct {
	server *http.Server
}

func (web *Web) Run(dataBase *db.DataBase) {
	web.server = &http.Server{Addr: ":3228"}

	handlers.GetDBObject(dataBase)
	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/result", handlers.ResultPage)

	log.Println("HTTP-Server is starting on \"http://localhost:3228\"...")
	err := web.server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}

func (web *Web) Finish() {
	err := web.server.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
}
