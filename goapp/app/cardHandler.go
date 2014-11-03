package app

import (
    "strconv"
    "strings"
	"lib/tool"
)

func CreateCard(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
    tool.VerifyParam(r, "name", tool.ParamNotNil())
    tool.VerifyParam(r, "action", tool.ParamNotNil())
    tool.VerifyParam(r, "actionContent", tool.ParamNotNil())
    tool.VerifyParam(r, "actionCost", tool.ParamNotNil())
    tool.VerifyParam(r, "actionType", tool.ParamNotNil())
    tool.VerifyParam(r, "atkType", tool.ParamNotNil())
    tool.VerifyParam(r, "content", tool.ParamNotNil())
    tool.VerifyParam(r, "tag", tool.ParamNotNil())
    tool.VerifyParam(r, "type", tool.ParamNotNil())
    tool.VerifyParam(r, "level", tool.ParamNotNil())
    tool.VerifyParam(r, "magic", tool.ParamNotNil())
    tool.VerifyParam(r, "power", tool.ParamNotNil())
    tool.VerifyParam(r, "weight", tool.ParamNotNil())
    
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
    
    card := CardEntity{
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
    var cr ICardDAO = GetApp().GetCardDAO(r)
	if shouldUpdate {
		key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
		cr.Update( sys, cr.GetKey(sys, key, nil), card )
		return tool.Success(key)
	}else{
		key := cr.Create(sys, cr.NewKey(sys, nil), card)
		return tool.Success(key.IntID())
	}
}

func QueryCard(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
    var cr ICardDAO = GetApp().GetCardDAO(r)
    
    keys := r.Form["key"]
    if keys != nil {
        var cards []CardEntity
        for _, k := range keys {
            if len(strings.TrimSpace(k)) == 0 {
                continue
            }
            ki, _ := strconv.ParseInt(k, 10, 64)
            card := cr.Read( sys, cr.GetKey(sys, ki, nil) ).(CardEntity)
            cards = append(cards, card)
        }
        return tool.Success(cards)
    }
	
	query := cr.NewQuery(sys)
	
	hasType := len( r.Form["type"] ) > 0 
	if hasType {
		for _, tp := range r.Form["type"] {
			query = query.Filter("Type = ", tp)
		}
	}
	
	hasWeight := len( r.Form["weight"] ) > 0 
	if hasWeight {
		for _, w := range r.Form["weight"] {
			wi, _ := strconv.Atoi(w)
			query = query.Filter("Weight =", wi)
		}
	}
	
	hasAtkType := len( r.Form["atkType"] ) > 0 
	if hasAtkType {
		for _, v := range r.Form["atkType"] {
			query = query.Filter("AtkType =", v)
		}
	}
	
	hasActionType := len( r.Form["actionType"] ) > 0 
	if hasActionType {
		for _, v := range r.Form["actionType"] {
			query = query.Filter("ActionType =", v)
		}
	}
	
	hasLevel := len( r.Form["level"] ) > 0 
	if hasLevel {
		for _, v := range r.Form["level"] {
			vi, _ := strconv.Atoi(v)
			query = query.Filter("Level =", vi)
		}
	}
	
	hasOffset := len( r.Form["offset"] ) > 0 
	if hasOffset {
		offset, _ := strconv.Atoi( r.Form["offset"][0] )
		query = query.Offset(offset)
	}
	
	hasLimit := len( r.Form["limit"] ) > 0 
	if hasLimit {
		limit, _ := strconv.Atoi( r.Form["limit"][0] )
		query = query.Limit(limit)
	}
	
	hasActionContent := len( r.Form["actionContent"] ) > 0 
	if hasActionContent {
		for _, v := range r.Form["actionContent"] {
			query = query.Filter("ActionContent >=", v)
		}
	}
    
    var cards []interface{} = cr.ReadAll(sys, query)
    return tool.Success(cards)
}

func CreateCardSuit(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
    tool.VerifyParam(r, "name", tool.ParamNotNil())
    tool.VerifyParam(r, "id[]", tool.ParamNotNil())
    tool.VerifyParam(r, "description", tool.ParamNotNil())
    
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
    
    var cr ICardSuitDAO = GetApp().GetCardSuitDAO(r)    
    key := cr.Create(sys, cr.NewKey(sys, nil), CardSuitEntity{Name:name, Description:description, CardIds:ids})
    return tool.Success(key)
}

func PrintCardSuit(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
    tool.VerifyParam(r, "cardSuitId", tool.ParamNotNil())
    tool.VerifyParam(r, "page", tool.ParamNotNil())
    
    cardSuitId, _ := strconv.ParseInt(r.Form["cardSuitId"][0], 10, 64)
    page, _ := strconv.Atoi(r.Form["page"][0])

    var cr ICardDAO = GetApp().GetCardDAO(r)
    var csr ICardSuitDAO = GetApp().GetCardSuitDAO(r)
    
    cardSuit := csr.Read(sys, csr.GetKey(sys, cardSuitId, nil)).(CardSuitEntity)

    var cards []CardEntity
    for idx, cardId := range cardSuit.CardIds {
        if idx >= page*9 && idx < page*9+9 {
            cards = append( cards, cr.Read(sys, cr.GetKey(sys, cardId, nil)).(CardEntity) )
        }
    }
	return tool.Success(cards)
}