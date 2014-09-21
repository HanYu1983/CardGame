package hello

import (
    "net/http"
    "appengine"
    "appengine/datastore"
)

type CardRepository struct {
    Kind string
    Request *http.Request
}

func (cr *CardRepository) Create(card CardPO) int64 {
    c := appengine.NewContext(cr.Request)
    key := datastore.NewIncompleteKey(c, cr.Kind, nil)
    retkey, err := datastore.Put(c, key, &card)
    if err != nil {
        panic(err.Error())
    }
    return retkey.IntID()
}

func (cr *CardRepository) Read(key int64) CardPO {
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    var card CardPO
    err := datastore.Get(c, pKey, &card)
    if err != nil {
        panic(err.Error())
    }
    card.Key = key
    return card
}

func (cr *CardRepository) Update(key int64, card CardPO) {
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    _, err := datastore.Put(c, pKey, card)
    if err != nil {
        panic(err.Error())
    }
}

func (cr *CardRepository) Delete(key int64) {
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    err := datastore.Delete(c, pKey)
    if err != nil {
        panic(err.Error())
    }
}

func (cr *CardRepository) GetAll() []CardPO {
    c := appengine.NewContext(cr.Request)
    q := datastore.NewQuery(cr.Kind)
    var cards []CardPO
    keys, err := q.GetAll(c, &cards)
    if err != nil {
        panic( err.Error() )
    }
    var retcards []CardPO
    for idx, card := range cards {
        card.Key = keys[idx].IntID()
        retcards = append(retcards, card)
    }
    return retcards
}