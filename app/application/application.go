package application

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogasab/bookstore_items-api/app/client/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapURLS()
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:5003",
	}

	log.Fatalln(server.ListenAndServe())
}
