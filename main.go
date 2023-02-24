package main

import (
	router "SuperDuper/routes"
	"log"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr:         ":9001",
		Handler:      router.Router(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is running, enter to the => http://localhost" + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
