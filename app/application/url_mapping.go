package application

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/handlers"
)

func mapURLS() {
	router.HandleFunc("/ping", handlers.NewPingHandler().Ping).Methods(http.MethodGet)
}
