package main

import (
	"go_STAN/internal/streaming"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//var cache *db.Cache = db.InitCache()
	var stanUsers streaming.StanUsers
	streaming.Run(&stanUsers)

	// CTRL+C - stop the app
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Stopping the app...")
	streaming.Finish(&stanUsers)
	log.Print("The app has stopped")
}
