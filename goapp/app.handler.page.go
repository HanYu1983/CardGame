package hello

import (
    "net/http"
    "html/template"
    "strconv"
)

type Friend struct {
    Fname string
}

type Person struct {
    UserName string
    Emails   []string
    Friends  []*Friend
}


func TestView(w http.ResponseWriter, r *http.Request) interface{} {
    f1 := Friend{Fname: "minux.ma"}
    f2 := Friend{Fname: "xushiwei"}
    t := template.New("fieldname example")
    t, _ = t.Parse(stringWithFile("tmpl/Test.html"))
    p := Person{UserName: "Astaxie",
        Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
        Friends: []*Friend{&f1, &f2}}
    
    w.Header().Set("Content-Type", "text/html")
    t.Execute(w, p)
    return CustomView
}


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
    
    http.Redirect(w, r, r.URL.Path+"?cmd=EditCardSuitPage&key="+strconv.FormatInt(key, 16), 301)
    return CustomView
}

type EditCardSuitPageModel struct{
    CardSuit CardSuitPO
    Cards []CardPO
}

func EditCardSuitPage(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "key", ParamNotNil())
    key, _ := strconv.ParseInt(r.Form["key"][0], 10, 64)
    
    var cr ICardRepository = GetApp().GetCardRepository(r)
    var csr ICardSuitRepository = GetApp().GetCardSuitRepository(r)
    
    cardSuit := csr.Read(key).(CardSuitPO)
    model := EditCardSuitPageModel{CardSuit: cardSuit}
    
    for _, cardId := range cardSuit.CardIds {
        model.Cards = append( model.Cards, cr.Read(cardId).(CardPO) )
    }
    
    w.Header().Set("Content-Type", "text/html")
    templateWithFile("query card", "tmpl/EditCardSuit.html").Execute(w, model)
    return CustomView
}

func ModifyCardWithCardSuit(w http.ResponseWriter, r *http.Request) interface{} {
    VerifyParam(r, "key", ParamNotNil())
    VerifyParam(r, "cardId", ParamNotNil())
    VerifyParam(r, "type", ParamNotNil())
    ty := r.Form["type"][0]
    switch ty {
    case "add":
        break
    case "delete":
        break
    }
    return CustomView
}