package hello

import (
    "net/http"
)

var app IApp = func()IApp{
	var ret SimpleApp
	ret.Init()
	return &ret
}()

func GetApp() IApp {
	return app
}

type SimpleApp struct {
	CardRepository CardRepository
	CardSuitRepository CardSuitRepository
}

func (app *SimpleApp) Init(){
	app.CardRepository.Init()
	app.CardSuitRepository.Init()
}

func (app *SimpleApp) GetCardRepository(request *http.Request) ICardRepository {
    return &app.CardRepository
}

func (app *SimpleApp) GetCardSuitRepository(request *http.Request) ICardSuitRepository {
    return &app.CardSuitRepository
}