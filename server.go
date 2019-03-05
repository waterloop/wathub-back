package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func baseFunc(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(nil)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", baseFunc)
    log.Fatal(http.ListenAndServe(":8000", r))
}
