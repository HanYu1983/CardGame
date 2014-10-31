package hello

import (
    "fmt"
    "net/http"
	"lib/tool"
	"app"
)

func init() {
    FuncFrontControl := func(w http.ResponseWriter, r *http.Request){
        tool.FrontController(w, r,
            tool.ActionMap{
                "CreateCard": app.CreateCard,
                "QueryCard": app.QueryCard,
                "CreateCardSuit": app.CreateCardSuit,
                "PrintCardSuit": app.PrintCardSuit,
            },
        )
    }
    
    PageFrontControl := func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Pragma", "No-cache")
        w.Header().Set("Cache-Control", "no-cache,no-store,max-age=0")
        w.Header().Set("Expires", "1")
        
        tool.FrontController(w, r,
            tool.ActionMap{
                "QueryCardPage": app.QueryCardPage,
                "AddCardSuit": app.AddCardSuit,
                "QueryCardSuitPage": app.QueryCardSuitPage,
                "EditCardSuitPage": app.EditCardSuitPage,
                "ModifyCardWithCardSuit": app.ModifyCardWithCardSuit,
            },
        )
    }
    
    http.HandleFunc("/", handler)
    http.HandleFunc("/Func", FuncFrontControl)
    http.HandleFunc("/Page", PageFrontControl)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}