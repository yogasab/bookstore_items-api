package items

import (
	"encoding/json"
	"errors"
	"strings"

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

func (i *Item) GetByID(ID string) rest_errors_utils.RestErrors {
	result, err := elasticsearch.ESClient.GetByID(indexItems, indexType, ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors_utils.NewNotFoundError(err.Error())
		}
		return rest_errors_utils.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	bytesResult, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors_utils.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	if err = json.Unmarshal(bytesResult, &i); err != nil {
		return rest_errors_utils.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	i.ID = ID
	return nil
}
