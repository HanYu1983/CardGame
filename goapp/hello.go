package hello

import (
    "fmt"
    "net/http"
	"lib/tool"
)

func init() {
    FuncFrontControl := func(w http.ResponseWriter, r *http.Request){
        tool.FrontController(w, r,
            tool.ActionMap{
                "CreateCard": CreateCard,
                "QueryCard": QueryCard,
                "CreateCardSuit": CreateCardSuit,
                "PrintCardSuit": PrintCardSuit,
            },
        )
    }
    
    PageFrontControl := func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Pragma", "No-cache")
        w.Header().Set("Cache-Control", "no-cache,no-store,max-age=0")
        w.Header().Set("Expires", "1")
        
        tool.FrontController(w, r,
            tool.ActionMap{
                "QueryCardPage": QueryCardPage,
                "AddCardSuit": AddCardSuit,
                "QueryCardSuitPage": QueryCardSuitPage,
                "EditCardSuitPage": EditCardSuitPage,
                "ModifyCardWithCardSuit": ModifyCardWithCardSuit,
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