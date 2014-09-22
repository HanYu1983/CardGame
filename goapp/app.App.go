package hello

import (
    "net/http"
    "appengine"
    "appengine/datastore"
)

var app App

func GetApp() *App{
    return &app
}

type App struct {
}

func (app *App) GetCardRepository(request *http.Request) IRepository {
    return &CardRepository{
                DefaultRepository{
                    Request: request, 
                    Kind: "Card",
                    PutFn: func(ctx appengine.Context, key *datastore.Key, po interface{}) int64 {
                        var card CardPO
                        retkey, err := datastore.Put(ctx, key, &card)
                        if err != nil {
                            panic(err.Error())
                        }
                        return retkey.IntID()
                    },
                    GetFn: func(ctx appengine.Context, key *datastore.Key) interface{} {
                        var card CardPO
                        err := datastore.Get(ctx, key, &card)
                        if err != nil {
                            panic(err.Error())
                        }
                        card.Key = key.IntID()
                        return card
                    },
                    GetAllFn: func(ctx appengine.Context, q *datastore.Query) []interface{} {
                        var cards []CardPO
                        keys, err := q.GetAll(ctx, &cards)
                        if err != nil {
                            panic( err.Error() )
                        }
                        var retcards []interface{}
                        for idx, card := range cards {
                            card.Key = keys[idx].IntID()
                            retcards = append(retcards, card)
                        }
                        return retcards
                    },
                },
            }
}