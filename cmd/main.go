package main

import (
	"go_STAN/internal/db"
	"go_STAN/internal/streaming"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("The app has started! Stop the app - CTRL+C")

	var dataBase *db.DataBase = db.InitDB()
	var stanUsers streaming.StanUsers
	streaming.Run(&stanUsers, dataBase)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Stopping the app...")

	streaming.Finish(&stanUsers)
	dataBase.Close()

	log.Println("The app has stopped!")
}
