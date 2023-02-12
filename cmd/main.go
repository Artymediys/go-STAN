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

	//streaming.Run(db)

	// CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Stopping the app...")
	streaming.Finish()
	log.Print("The app has stopped")
}
