package services

import (
	"github.com/yogasab/bookstore_items-api/app/domain/es_queries"
	"github.com/yogasab/bookstore_items-api/app/domain/items"
	rest_errors "github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

type ItemsService interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErrors)
	Get(ID string) (*items.Item, rest_errors.RestErrors)
	Search(query es_queries.ESQueries) ([]items.Item, rest_errors.RestErrors)
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

func (s *itemService) Get(ID string) (*items.Item, rest_errors.RestErrors) {
	item := items.Item{}
	if err := item.GetByID(ID); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query es_queries.ESQueries) ([]items.Item, rest_errors.RestErrors) {
	item := items.Item{}
	items, err := item.Search(query)
	if err != nil {
		return nil, err
	}
	return *items, nil
}
