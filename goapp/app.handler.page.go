package hello

import (
    "net/http"
    "strconv"
    "sort"
)

type QueryCardPageModel struct {
    Cards []interface{}
}

func QueryCardPage(w http.ResponseWriter, r *http.Request) interface{} {
    var cr ICardRepository = GetApp().GetCardRepository(r)
    model := QueryCardPageModel{Cards: cr.GetAll()}
    
    w.Header().Set("Content-Type", "text/html")
    templateWithFile("query card", "tmpl/QueryCard.html").Execute(w, model)
    return CustomView
}

type QueryCardSuitPageModel struct {
    CardSuits []interface{}
}

func QueryCardSuitPage(w http.ResponseWriter, r *http.Request) interface{} {
    var cr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    model := QueryCardSuitPageModel{CardSuits: cr.GetAll()}
    
    w.Header().Set("Content-Type", "text/html")
    templateWithFile("query card", "tmpl/QueryCardSuit.html").Execute(w, model)
    return CustomView
}

func AddCardSuit(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "name", ParamNotNil())
    name := r.Form["name"][0]
    
    var cr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    key := cr.Create(CardSuitPO{Name:name})
    
    http.Redirect(w, r, r.URL.Path+"?cmd=EditCardSuitPage&key="+strconv.FormatInt(key, 10), 302)
    return CustomView
}

type EditCardSuitPageModel struct{
    CardSuit CardSuitPO
    Cards []CardPO
    AllCards []interface{}
}

func EditCardSuitPage(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "key", ParamNotNil())
    key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
    
    var cr ICardRepository = GetApp().GetCardRepository(r)
    var csr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    
    cardSuit := csr.Read(key).(CardSuitPO)
    model := EditCardSuitPageModel{CardSuit: cardSuit}
    model.AllCards = cr.GetAll()
    
    sort.Sort(int64Array(cardSuit.CardIds))
    
    for _, cardId := range cardSuit.CardIds {
        model.Cards = append( model.Cards, cr.Read(cardId).(CardPO) )
    }
    
    w.Header().Set("Content-Type", "text/html")
    templateWithFile("query card", "tmpl/EditCardSuit.html").Execute(w, model)
    return CustomView
}

func ModifyCardWithCardSuit(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "key", ParamNotNil())
    VerifyParam(r, "type", ParamNotNil())
    
    ty := r.Form["type"][0]
    key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
    
    var csr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    cardSuit := csr.Read(key).(CardSuitPO)
    
    switch ty {
    case "none":
        VerifyParam(r, "name", ParamNotNil())
        VerifyParam(r, "description", ParamNotNil())
        
        name := r.Form["name"][0]
        description := r.Form["description"][0]
        cardSuit.Name = name
        cardSuit.Description = description
        csr.Update(key, cardSuit)
        break
        
    case "add":
        VerifyParam(r, "cardId", ParamNotNil())
        cardId, _ := strconv.ParseInt(r.Form["cardId"][0], 10, 64) 
        cardSuit.CardIds = append(cardSuit.CardIds, cardId)
        csr.Update(key, cardSuit)
        break
        
    case "delete":
        VerifyParam(r, "cardId", ParamNotNil())
        cardId, _ := strconv.ParseInt(r.Form["cardId"][0], 10, 64) 
        for idx, v := range cardSuit.CardIds {
            if v == cardId {
                cardSuit.CardIds = append( cardSuit.CardIds[:idx], cardSuit.CardIds[(idx+1):]... )
                csr.Update(key, cardSuit)
                break
            }
        }
        break
    }
    http.Redirect(w, r, r.URL.Path+"?cmd=EditCardSuitPage&key="+strconv.FormatInt(key, 10), 302)
    return CustomView
}

