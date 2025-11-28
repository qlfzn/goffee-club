package products

import (
    "encoding/json"
    "os"

    repo "github.com/qlfzn/goffee-club/internal/repository"
)

type Service interface {
    GetProducts() ([]repo.MenuItem, error)
}

type svc struct{
    // define dependencies
}

func NewService() Service{
    return &svc{}
}

func (s *svc) GetProducts() ([]repo.MenuItem, error) {
    file, err := os.ReadFile("menu.json")
    if err != nil {
        return nil, err
    }

    var menu repo.Menu
    err = json.Unmarshal(file, &menu)
    if err != nil {
        return nil, err
    }

    return menu.MenuItems, nil
}