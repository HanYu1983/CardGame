package hello

import (
    "fmt"
    "net/http"
)

func init() {
    FuncFrontControl := func(w http.ResponseWriter, r *http.Request){
        FrontControl(w, r,
            ActionMap{
                "CreateCard": CreateCard,
                "QueryCard": QueryCard,
                
                "CreateCardSuit": CreateCardSuit,
            },
        )
    }
    
    PageFrontControl := func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Pragma", "No-cache")
        w.Header().Set("Cache-Control", "no-cache,no-store,max-age=0")
        w.Header().Set("Expires", "1")
        
        FrontControl(w, r,
            ActionMap{
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