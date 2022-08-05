package handlers

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/domain/items"
	"github.com/yogasab/bookstore_items-api/app/services"
	"github.com/yogasab/bookstore_oauth-go/oauth"
)

type ItemsHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsHandler struct {
}

func NewItemsHandler() ItemsHandler {
	return &itemsHandler{}
}

func (h *itemsHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	seller := oauth.GetCallerID(r)
	item := items.Item{
		Seller: seller,
	}

	_, err := services.NewItemsService().Create(item)
	if err != nil {
		return
	}
}
func (h *itemsHandler) Get(w http.ResponseWriter, r *http.Request) {

}
