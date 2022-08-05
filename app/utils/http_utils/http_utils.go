package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/yogasab/bookstore_items-api/app/utils/rest_errors_utils"
)

func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func ResponseJSONError(w http.ResponseWriter, err rest_errors_utils.RestErrors) {
	ResponseJSON(w, err.Code, err)
}
