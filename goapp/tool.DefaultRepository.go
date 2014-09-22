package hello

import (
    "net/http"
    "appengine"
    "appengine/datastore"
)

type DefaultRepository struct {
    Kind string
    Request *http.Request
    PutFn func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error)
    GetFn func(ctx appengine.Context, key *datastore.Key) (ret interface{}, err error)
    GetAllFn func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error )
}

func (cr *DefaultRepository) Create(po interface{}) int64 {
    c := appengine.NewContext(cr.Request)
    key := datastore.NewIncompleteKey(c, cr.Kind, nil)    
    retkey, err := cr.PutFn(c, key, po)
    if err != nil {
        panic(err.Error())
    }
    return retkey.IntID()
}

func (cr *DefaultRepository) Update(key int64, po interface{}){
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    cr.PutFn(c, pKey, po)
}
func (cr *DefaultRepository) Read(key int64) interface{}{
    c := appengine.NewContext(cr.Request)
    pKey := datastore.NewKey(c, cr.Kind, "", key, nil)
    ret, err := cr.GetFn(c, pKey)
    if err != nil {
        panic(err.Error())
    }
    return ret
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
    pos, _, err := cr.GetAllFn( c, q )
    if err != nil {
        panic( err.Error() )
    }
    return pos
}