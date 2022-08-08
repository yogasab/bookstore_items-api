package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/yogasab/bookstore_items-api/app/domain/es_queries"
	"github.com/yogasab/bookstore_items-api/app/domain/items"
	"github.com/yogasab/bookstore_items-api/app/services"
	"github.com/yogasab/bookstore_items-api/app/utils/http_utils"
	"github.com/yogasab/bookstore_items-api/app/utils/oauth"
	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

type ItemsHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsHandler struct {
	itemService services.ItemsService
}

func NewItemsHandler(itemService services.ItemsService) ItemsHandler {
	return &itemsHandler{itemService: itemService}
}

func (h *itemsHandler) Create(w http.ResponseWriter, r *http.Request) {
	errRest := oauth.AuthenticateRequest(r)
	if errRest != nil {
		http_utils.ResponseJSON(w, errRest.Code, errRest)
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
		http_utils.ResponseJSONError(w, restError)
		return
	}
	defer r.Body.Close()

	var item items.Item
	if err = json.Unmarshal(bytesResult, &item); err != nil {
		restError := rest_errors_utils.NewBadRequestError("invalid request item struct body")
		http_utils.ResponseJSONError(w, restError)
		return
	}

	item.Seller = sellerID

	result, errCreate := h.itemService.Create(item)
	if err != nil {
		http_utils.ResponseJSONError(w, errCreate)
		return
	}
	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (h *itemsHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])
	result, err := h.itemService.Get(itemID)
	if err != nil {
		http_utils.ResponseJSON(w, err.Code(), err)
		return
	}
	http_utils.ResponseJSON(w, http.StatusOK, result)
}

func (h *itemsHandler) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errRest := rest_errors_utils.NewBadRequestError("invalid json body")
		http_utils.ResponseJSONError(w, errRest)
		return
	}
	defer r.Body.Close()

	var query es_queries.ESQueries
	if err = json.Unmarshal(bytes, &query); err != nil {
		errRest := rest_errors_utils.NewBadRequestError("invalid json body")
		http_utils.ResponseJSONError(w, errRest)
		return
	}

	items, errRest := h.itemService.Search(query)
	if errRest != nil {
		errRest := rest_errors_utils.NewBadRequestError("invalid json body")
		http_utils.ResponseJSONError(w, errRest)
		return
	}

	http_utils.ResponseJSON(w, http.StatusOK, items)
}
