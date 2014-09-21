package hello

import (
    "net/http"
    "strconv"
)

func CreateCard(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "name", ParamNotNil())
    VerifyParam(r, "action", ParamNotNil())
    VerifyParam(r, "actionContent", ParamNotNil())
    VerifyParam(r, "actionCost", ParamNotNil())
    VerifyParam(r, "actionType", ParamNotNil())
    VerifyParam(r, "atkType", ParamNotNil())
    VerifyParam(r, "content", ParamNotNil())
    VerifyParam(r, "tag", ParamNotNil())
    VerifyParam(r, "type", ParamNotNil())
    VerifyParam(r, "level", ParamNotNil())
    VerifyParam(r, "magic", ParamNotNil())
    VerifyParam(r, "power", ParamNotNil())
    VerifyParam(r, "weight", ParamNotNil())
    
    name := r.Form["name"][0]
    action := r.Form["action"][0]
    actionContent := r.Form["actionContent"][0]
    actionCost := r.Form["actionCost"][0]
    actionType := r.Form["actionType"][0]
    atkType := r.Form["atkType"][0]
    content := r.Form["content"][0]
    tag := r.Form["tag"][0]
    tp := r.Form["type"][0]
    level, _ := strconv.Atoi(r.Form["level"][0])
    magic, _ := strconv.Atoi(r.Form["magic"][0])
    power, _ := strconv.Atoi(r.Form["power"][0])
    weight, _ := strconv.Atoi(r.Form["weight"][0])
    
    card := CardPO{
        Name: name,
        Action: action,
        ActionContent: actionContent, 
        ActionCost: actionCost,
        ActionType: actionType, 
        AtkType: atkType,
        Content: content,
        Level: level, 
        Magic: magic,
        Power: power,
        Tag: tag,
        Type: tp, 
        Weight: weight,
    }
    var cr ICardRepository = GetApp().GetCardRepository(r)
    key := cr.Create(card)
    return Success(key)
}

func QueryCard(w http.ResponseWriter, r *http.Request) interface{} {
    var cr ICardRepository = GetApp().GetCardRepository(r)
    var cards []CardPO = cr.GetAll()
    return Success(cards)
}