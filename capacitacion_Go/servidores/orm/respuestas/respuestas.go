package respuestas

import (
	"encoding/json"
	"net/http"
)

func ResponderJSON(obj interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(obj)
}
