package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/domain/items"
	"github.com/yogasab/bookstore_items-api/app/services"
	"github.com/yogasab/bookstore_items-api/app/utils/http_utils"
	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
	"github.com/yogasab/bookstore_oauth-go/oauth"
)

type ItemsHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsHandler struct {
	itemService services.ItemsService
}

func NewItemsHandler(itemService services.ItemsService) ItemsHandler {
	return &itemsHandler{itemService: itemService}
}

func (h *itemsHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseJSON(w, err.Code, err)
		return
	}
	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restError := rest_errors_utils.NewUnauthorizedError("invalid access token")
		http_utils.ResponseJSON(w, http.StatusUnauthorized, restError)
		return
	}

	bytesResult, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restError := rest_errors_utils.NewBadRequestError("invalid request body")
		http_utils.ResponseJSONError(w, *restError)
		return
	}
	defer r.Body.Close()

	var item items.Item
	if err = json.Unmarshal(bytesResult, &item); err != nil {
		restError := rest_errors_utils.NewBadRequestError("invalid request item struct body")
		http_utils.ResponseJSONError(w, *restError)
		return
	}

	item.Seller = sellerID

	result, errCreate := h.itemService.Create(item)
	if err != nil {
		http_utils.ResponseJSONError(w, *errCreate)
		return
	}
	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (h *itemsHandler) Get(w http.ResponseWriter, r *http.Request) {
}
