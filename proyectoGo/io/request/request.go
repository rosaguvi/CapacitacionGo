package request 

import (
	"net/http"
	"encoding/json"
)

func Json(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(&obj)
}