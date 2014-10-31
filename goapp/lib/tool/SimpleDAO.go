package tool

import (
    "appengine"
    "appengine/datastore"
)

type SimpleDAO struct {
    Kind string
    PutFn func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error)
    GetFn func(ctx appengine.Context, keys []*datastore.Key) (ret []interface{}, err error)
    GetAllFn func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error )
}

func (cr *SimpleDAO) NewKey(sys ISystem, parent *datastore.Key) *datastore.Key {
	c := appengine.NewContext(sys.GetRequest())
	return datastore.NewIncompleteKey(c, cr.Kind, parent) 
}

func (cr *SimpleDAO) GetKey(sys ISystem, key int64, parent *datastore.Key) *datastore.Key{
	c := appengine.NewContext(sys.GetRequest())
	return datastore.NewKey(c, cr.Kind, "", key, parent) 
}

func (cr *SimpleDAO) NewQuery(sys ISystem) *datastore.Query {
	return datastore.NewQuery(cr.Kind)
}

func (cr *SimpleDAO) Create(sys ISystem, key *datastore.Key, po interface{}) *datastore.Key {
    c := appengine.NewContext(sys.GetRequest())  
    retkey, err := cr.PutFn(c, key, po)
    if err != nil {
        panic(err.Error())
    }
    return retkey
}

func (cr *SimpleDAO) Update(sys ISystem, key *datastore.Key, po interface{}){
    c := appengine.NewContext(sys.GetRequest())
    cr.PutFn(c, key, po)
}
func (cr *SimpleDAO) Read(sys ISystem, key *datastore.Key) interface{}{
    ret := cr.ReadMulti(sys, []*datastore.Key{key})
	if len(ret) > 0{
		return ret[0]
	}else{
		return nil
	}
}
func (cr *SimpleDAO) ReadMulti(sys ISystem, keys []*datastore.Key) []interface{}{
    c := appengine.NewContext(sys.GetRequest())
    ret, err := cr.GetFn(c, keys)
    if err != nil {
        panic(err.Error())
    }
    return ret
}
func (cr *SimpleDAO) Delete(sys ISystem, key *datastore.Key){
    c := appengine.NewContext(sys.GetRequest())
    err := datastore.Delete(c, key)
    if err != nil {
        panic(err.Error())
    }
}
func (cr *SimpleDAO) ReadAll(sys ISystem, query *datastore.Query) []interface{}{
    c := appengine.NewContext(sys.GetRequest())
    pos, _, err := cr.GetAllFn( c, query )
    if err != nil {
        panic( err.Error() )
    }
    return pos
}