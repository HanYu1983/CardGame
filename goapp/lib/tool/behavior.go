package tool

import (
	"net/http"
	"appengine/datastore"
)

type IDataAccessObject interface {
	NewKey(sys ISystem, parent *datastore.Key) *datastore.Key 
	GetKey(sys ISystem, key int64, parent *datastore.Key) *datastore.Key
	NewQuery(sys ISystem) *datastore.Query
	
    Create(sys ISystem, key *datastore.Key, po interface{}) *datastore.Key
    Update(sys ISystem, key *datastore.Key, po interface{})
    Read(sys ISystem, key *datastore.Key) interface{}
	ReadMulti(sys ISystem, key []*datastore.Key) []interface{}
	ReadAll(sys ISystem, query *datastore.Query) []interface{}
    Delete(sys ISystem, key *datastore.Key)
}

type ICookieManager interface {
    GetValue(sys ISystem)(bool, string)
    SetValue(sys ISystem, v string)
    Clear(sys ISystem)
}

type ILogger interface {
	Log(msg interface{})
}

type ISystem interface {
	GetRequest() *http.Request
	GetResponse() http.ResponseWriter
	Log(msg interface{})
}

type IEventManager interface{
	AddSender(key string, fn func(mgr IEventManager))
	RemoveSender(key string, fn func(mgr IEventManager))
	AddReceiver(key string, fn interface{})
	RemoveReceiver(key string, fn interface{})
	SendAgain()
	Send(key string, params ... interface{})
}