package models

type Member struct {
    Name     string   `json:"name"`
    Email    string   `json:"email"`
    Title    string   `json:"title"`
    Lead     bool     `json:"lead"`
    Owner    bool     `json:"owner"`
    Access   []string `json:"access"`
    Teams    []string `json:"teams"`
    Projects []string `json:"projects"`
}
