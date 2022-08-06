package items

import (
	"github.com/yogasab/bookstore_items-api/app/client/elasticsearch"
	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

var (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors_utils.RestErrors {
	response, err := elasticsearch.ESClient.Index(indexItems, i)
	if err != nil {
		rest_errors_utils.NewInternalServerError(err.Error())
	}
	i.ID = response.Id
	return nil
}
