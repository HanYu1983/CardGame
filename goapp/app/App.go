package app

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
	CardDAO CardDAO
	CardSuitDAO CardSuitDAO
}

func (app *SimpleApp) Init(){
	app.CardDAO.Init()
	app.CardSuitDAO.Init()
}

func (app *SimpleApp) GetCardDAO(request *http.Request) ICardDAO {
    return &app.CardDAO
}

func (app *SimpleApp) GetCardSuitDAO(request *http.Request) ICardSuitDAO {
    return &app.CardSuitDAO
}