package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "wathub-back/models"
    "wathub-back/utils"
)

func baseFunc(w http.ResponseWriter, r *http.Request) {
     utils.JsonRespond(w, models.Term{})
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", baseFunc)
    log.Fatal(http.ListenAndServe(":8000", r))
}
