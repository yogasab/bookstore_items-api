package services

import (
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/domain/items"
	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
	rest_errors "github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

type ItemsService interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErrors)
	Get(string) (*items.Item, rest_errors.RestErrors)
}

type itemService struct {
}

func NewItemsService() ItemsService {
	return &itemService{}
}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErrors) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *itemService) Get(string) (*items.Item, rest_errors.RestErrors) {
	return nil, rest_errors_utils.NewRestErrors(http.StatusNotImplemented, "failed", "not implemented", nil)
}
