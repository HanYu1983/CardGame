package hello

import (
    "net/http"
    "strconv"
    "strings"
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
    
	shouldUpdate := len(r.Form["key"]) > 0
	
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
	if shouldUpdate {
		key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
		cr.Update( key, card )
		return Success(key)
	}else{
		key := cr.Create(card)
		return Success(key)
	}
}

func QueryCard(w http.ResponseWriter, r *http.Request) interface{} {
    var cr ICardRepository = GetApp().GetCardRepository(r)
    
    keys := r.Form["key"]
    if keys != nil {
        var cards []CardPO
        for _, k := range keys {
            if len(strings.TrimSpace(k)) == 0 {
                continue
            }
            ki, _ := strconv.ParseInt(k, 10, 64)
            card := cr.Read( ki ).(CardPO)
            cards = append(cards, card)
        }
        return Success(cards)
    }
    
    var cards []interface{} = cr.GetAll()
    return Success(cards)
}

func CreateCardSuit(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "name", ParamNotNil())
    VerifyParam(r, "id[]", ParamNotNil())
    VerifyParam(r, "description", ParamNotNil())
    
    name := r.Form["name"][0]
    description := r.Form["description"][0]
    idstr := r.Form["id[]"]
    
    var ids []int64
    for _, v := range idstr {
        id, err := strconv.ParseInt(v, 10, 64)
        if err != nil { 
            panic(err.Error())
        }
        ids = append( ids, id )
    }
    
    var cr ICardSuitRepository = GetApp().GetCardSuitRepository(r)    
    key := cr.Create(CardSuitPO{Name:name, Description:description, CardIds:ids})
    return Success(key)
}

func PrintCardSuit(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "cardSuitId", ParamNotNil())
    VerifyParam(r, "page", ParamNotNil())
    
    cardSuitId, _ := strconv.ParseInt(r.Form["cardSuitId"][0], 10, 64)
    page, _ := strconv.Atoi(r.Form["page"][0])
    
    var cr ICardRepository = GetApp().GetCardRepository(r)
    var csr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    
    cardSuit := csr.Read(cardSuitId).(CardSuitPO)
    
    var cards []CardPO
    for idx, cardId := range cardSuit.CardIds {
        if idx >= page*9 && idx < page*9+9 {
            cards = append( cards, cr.Read(cardId).(CardPO) )
        }
    }
    return Success(cards)
}