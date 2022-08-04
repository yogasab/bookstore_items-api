package application

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapURLS()
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:5000",
	}

	log.Fatalln(server.ListenAndServe())
}
