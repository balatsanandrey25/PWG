package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, w http.ResponseWriter, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
	}

}
