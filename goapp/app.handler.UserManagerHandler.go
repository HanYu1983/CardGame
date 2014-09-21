package hello

import (
    "net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "name", ParamNotNil())
    name := r.Form["name"][0]
    card := CardPO{Name:name}
    
    var cr ICardRepository = GetApp().GetCardRepository(r)
    cr.Create(card)
    return DefaultResult{Success: true, Info:card}
}