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
        FrontControl(w, r,
            ActionMap{
                "TestView": TestView,
                "QueryCardPage": QueryCardPage,
                "AddCardSuit": AddCardSuit,
                "QueryCardSuitPage": QueryCardSuitPage,
                "EditCardSuitPage": EditCardSuitPage,
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