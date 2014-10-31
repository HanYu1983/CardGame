package app

import (
	"lib/tool"
	"appengine"
    "appengine/datastore"
)

type CardSuitDAO struct {
    tool.SimpleDAO
}

func (r *CardSuitDAO) Init(){
	r.Kind = "CardSuit"
	r.PutFn = func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
		card := po.(CardSuitEntity)
		return  datastore.Put(ctx, key, &card)
	}
	r.GetFn = func(ctx appengine.Context, keys []*datastore.Key) (ret []interface{}, err error) {
		var entities []CardSuitEntity
		entities = make([]CardSuitEntity, len(keys))
		err = datastore.GetMulti(ctx, keys, entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
	r.GetAllFn = func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
		var entities []CardSuitEntity
		keys, err = q.GetAll(ctx, &entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
}