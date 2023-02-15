package main

import (
	"go_STAN/internal/db"
	"go_STAN/internal/streaming"
	"go_STAN/internal/web"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var dataBase *db.DataBase = db.InitDB()
	var stanUsers streaming.StanUsers
	var server web.Web

	streaming.Run(&stanUsers, dataBase)
	go server.Run(dataBase)

	log.Println("The app has started! Stop the app - CTRL+C")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Stopping the app...")
	streaming.Finish(&stanUsers)
	dataBase.Close()
	server.Finish()
	log.Println("The app has stopped!")
}
