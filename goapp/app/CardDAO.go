package app

import (
	"lib/tool"
	"appengine"
    "appengine/datastore"
)

type CardDAO struct {
    tool.SimpleDAO
}

func (r *CardDAO) Init(){
	r.Kind = "Card"
	r.PutFn = func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
		card := po.(CardEntity)
		return  datastore.Put(ctx, key, &card)
	}
	r.GetFn = func(ctx appengine.Context, keys []*datastore.Key) (ret []interface{}, err error) {
		var entities []CardEntity
		entities = make([]CardEntity, len(keys))
		err = datastore.GetMulti(ctx, keys, entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
	r.GetAllFn = func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
		q = q.Order("Type")
		var entities []CardEntity
		keys, err = q.GetAll(ctx, &entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
}