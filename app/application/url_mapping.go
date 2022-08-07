package application

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/handlers"
	"github.com/yogasab/bookstore_items-api/app/services"
)

func mapURLS() {
	itemsHandler := handlers.NewItemsHandler(services.NewItemsService())

	router.HandleFunc("/ping", handlers.NewPingHandler().Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", itemsHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", itemsHandler.Get).Methods(http.MethodGet)
}
