package hello

import (
    "net/http"
    "appengine"
    "appengine/datastore"
)

type DefaultRepository struct {
    Kind string
    Request *http.Request
    PutFn func(ctx appengine.Context, key *datastore.Key, po interface{}) int64
    GetFn func(ctx appengine.Context, key *datastore.Key) interface{}
    GetAllFn func(ctx appengine.Context, q *datastore.Query) []interface{}
}

func (cr *DefaultRepository) Create(po interface{}) int64 {
    c := appengine.NewContext(cr.Request)
    key := datastore.NewIncompleteKey(c, cr.Kind, nil)
    return cr.PutFn(c, key, po)
}

func (cr *DefaultRepository) Update(key int64, po interface{}){
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    cr.PutFn(c, pKey, po)
}
func (cr *DefaultRepository) Read(key int64) interface{}{
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    return cr.GetFn(c, pKey)
}
func (cr *DefaultRepository) Delete(key int64){
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    err := datastore.Delete(c, pKey)
    if err != nil {
        panic(err.Error())
    }
}
func (cr *DefaultRepository) GetAll() []interface{}{
    c := appengine.NewContext(cr.Request)
    q := datastore.NewQuery(cr.Kind)
    return cr.GetAllFn( c, q )
}