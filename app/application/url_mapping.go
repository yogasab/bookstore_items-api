package application

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/handlers"
)

func mapURLS() {
	itemsHandler := handlers.NewItemsHandler()
	router.HandleFunc("/ping", handlers.NewPingHandler().Ping).Methods(http.MethodGet)
	router.HandleFunc("/get", itemsHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/create", itemsHandler.Create).Methods(http.MethodPost)
}
