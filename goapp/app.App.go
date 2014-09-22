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

func (app *App) GetCardRepository(request *http.Request) ICardRepository {
    return &CardRepository{
                DefaultRepository{
                    Request: request, 
                    Kind: "Card",
                    PutFn: func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
                        card := po.(CardPO)
                        return  datastore.Put(ctx, key, &card)
                    },
                    GetFn: func(ctx appengine.Context, key *datastore.Key) (ret interface{}, err error) {
                        var card CardPO
                        err = datastore.Get(ctx, key, &card)
                        card.Key = key.IntID()
                        ret = card
                        return
                    },
                    GetAllFn: func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
                        var cards []CardPO
                        keys, err = q.GetAll(ctx, &cards)
                        for idx, card := range cards {
                            card.Key = keys[idx].IntID()
                            ret = append(ret, card)
                        }
                        return
                    },
                },
            }
}

func (app *App) GetCardSuitRepository(request *http.Request) ICardSuitRepository {
    return &CardSuitRepository{
                DefaultRepository{
                    Request: request, 
                    Kind: "CardSuit",
                    PutFn: func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
                        card := po.(CardSuitPO)
                        return  datastore.Put(ctx, key, &card)
                    },
                    GetFn: func(ctx appengine.Context, key *datastore.Key) (ret interface{}, err error) {
                        var card CardSuitPO
                        err = datastore.Get(ctx, key, &card)
                        card.Key = key.IntID()
                        ret = card
                        return
                    },
                    GetAllFn: func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
                        var cards []CardSuitPO
                        keys, err = q.GetAll(ctx, &cards)
                        for idx, card := range cards {
                            card.Key = keys[idx].IntID()
                            ret = append(ret, card)
                        }
                        return
                    },
                },
            }
}