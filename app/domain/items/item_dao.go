package items

import (
	"errors"

	"github.com/yogasab/bookstore_items-api/app/client/elasticsearch"
	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

var (
	indexItems = "items"
	indexType  = "_doc"
)

func (i *Item) Save() rest_errors_utils.RestErrors {
	response, err := elasticsearch.ESClient.Index(indexItems, indexType, i)
	if err != nil {
		rest_errors_utils.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	i.ID = response.Id
	return nil
}
