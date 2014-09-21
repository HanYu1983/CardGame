package hello

import (
    "fmt"
    "net/http"
)

func init() {
    FrontControl := func(w http.ResponseWriter, r *http.Request){
        FrontControl(w, r,
            ActionMap{
                "CreateCard": CreateCard,
            },
        )
    }
    
    http.HandleFunc("/", handler)
    http.HandleFunc("/test", FrontControl)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}