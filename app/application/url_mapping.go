package application

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/handlers"
	"github.com/yogasab/bookstore_items-api/app/services"
)

func mapURLS() {
	itemsHandler := handlers.NewItemsHandler(services.NewItemsService())

	router.HandleFunc("/ping", handlers.NewPingHandler().Ping).Methods(http.MethodGet)

	router.HandleFunc("/items/get", itemsHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/create", itemsHandler.Create).Methods(http.MethodPost)
}
