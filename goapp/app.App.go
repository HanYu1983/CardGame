package hello

import (
    "net/http"
)

var app App

func GetApp() *App{
    return &app
}

type App struct {
}

func (app *App) GetCardRepository(request *http.Request) ICardRepository {
    return &CardRepository{Request: request, Kind: "Card"}
}