package app

import (
    "net/http"
    "strconv"
    "sort"
	"lib/tool"
)

type QueryCardPageModel struct {
    Cards []interface{}
}

func QueryCardPage(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
	w := sys.GetResponse()
    var cr ICardDAO = GetApp().GetCardDAO(r)
    model := QueryCardPageModel{Cards: cr.ReadAll(sys, cr.NewQuery(sys))}
    
    w.Header().Set("Content-Type", "text/html")
    tool.TemplateWithFile("query card", "tmpl/QueryCard.html").Execute(w, model)
    return tool.CustomView
}

type QueryCardSuitPageModel struct {
    CardSuits []interface{}
}

func QueryCardSuitPage(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
	w := sys.GetResponse()
    var cr ICardSuitDAO = GetApp().GetCardSuitDAO(r)
    model := QueryCardSuitPageModel{CardSuits: cr.ReadAll(sys, cr.NewQuery(sys))}
    
    w.Header().Set("Content-Type", "text/html")
    tool.TemplateWithFile("query card", "tmpl/QueryCardSuit.html").Execute(w, model)
    return tool.CustomView
}

func AddCardSuit(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
	w := sys.GetResponse()
    tool.VerifyParam(r, "name", tool.ParamNotNil())
    name := r.Form["name"][0]
    
    var cr ICardSuitDAO = GetApp().GetCardSuitDAO(r)
    key := cr.Create(sys, cr.NewKey(sys, nil), CardSuitEntity{Name:name})
    
    http.Redirect(w, r, r.URL.Path+"?cmd=EditCardSuitPage&key="+strconv.FormatInt(key.IntID(), 10), 302)
    return tool.CustomView
}

type EditCardSuitPageModel struct{
    CardSuit CardSuitEntity
    Cards []CardEntity
    AllCards []interface{}
}

func EditCardSuitPage(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
	w := sys.GetResponse()
    tool.VerifyParam(r, "key", tool.ParamNotNil())
    key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
    
    var cr ICardDAO = GetApp().GetCardDAO(r)
    var csr ICardSuitDAO = GetApp().GetCardSuitDAO(r)
    
    cardSuit := csr.Read(sys, csr.GetKey(sys, key, nil)).(CardSuitEntity)
    model := EditCardSuitPageModel{CardSuit: cardSuit}
    model.AllCards = cr.ReadAll(sys, cr.NewQuery(sys))
    
    sort.Sort(tool.ByTypeInt64(cardSuit.CardIds))
    
    for _, cardId := range cardSuit.CardIds {
        model.Cards = append( model.Cards, cr.Read(sys, cr.GetKey(sys, cardId, nil)).(CardEntity) )
    }
    
    w.Header().Set("Content-Type", "text/html")
    tool.TemplateWithFile("query card", "tmpl/EditCardSuit.html").Execute(w, model)
    return tool.CustomView
}

func ModifyCardWithCardSuit(sys tool.ISystem) interface{} {
	r := sys.GetRequest()
	w := sys.GetResponse()
    tool.VerifyParam(r, "key", tool.ParamNotNil())
    tool.VerifyParam(r, "type", tool.ParamNotNil())
    
    ty := r.Form["type"][0]
    key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
    
    var csr ICardSuitDAO = GetApp().GetCardSuitDAO(r)
    cardSuit := csr.Read(sys, csr.GetKey(sys, key, nil)).(CardSuitEntity)
    
    switch ty {
    case "none":
        tool.VerifyParam(r, "name", tool.ParamNotNil())
        tool.VerifyParam(r, "description", tool.ParamNotNil())
        
        name := r.Form["name"][0]
        description := r.Form["description"][0]
        cardSuit.Name = name
        cardSuit.Description = description
        csr.Update(sys, csr.GetKey(sys, key, nil), cardSuit)
        break
        
    case "add":
        tool.VerifyParam(r, "cardId", tool.ParamNotNil())
        cardId, _ := strconv.ParseInt(r.Form["cardId"][0], 10, 64) 
        cardSuit.CardIds = append(cardSuit.CardIds, cardId)
        csr.Update(sys, csr.GetKey(sys, key, nil), cardSuit)
        break
        
    case "delete":
        tool.VerifyParam(r, "cardId", tool.ParamNotNil())
        cardId, _ := strconv.ParseInt(r.Form["cardId"][0], 10, 64) 
        for idx, v := range cardSuit.CardIds {
            if v == cardId {
                cardSuit.CardIds = append( cardSuit.CardIds[:idx], cardSuit.CardIds[(idx+1):]... )
                csr.Update(sys, csr.GetKey(sys, key, nil), cardSuit)
                break
            }
        }
        break
    }
    http.Redirect(w, r, r.URL.Path+"?cmd=EditCardSuitPage&key="+strconv.FormatInt(key, 10), 302)
    return tool.CustomView
}

