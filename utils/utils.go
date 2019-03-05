package utils

import (
    "encoding/json"
    "net/http"
)

func JsonRespond(w http.ResponseWriter, data interface{}) {
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
