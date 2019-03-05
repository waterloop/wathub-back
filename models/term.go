package models

type Term struct {
    Name    string   `json:"name"`
    Members []Member `json:"members"`
}
