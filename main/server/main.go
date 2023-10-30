package main

import (
	"gaspar44/TQS/controller"
	"log"
)

func main() {
	httpFileServer := controller.NewServer()

	if err := httpFileServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
