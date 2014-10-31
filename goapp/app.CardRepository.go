package hello

import (
	"lib/tool"
	"appengine"
    "appengine/datastore"
)

type CardRepository struct {
    tool.SimpleDAO
}

func (r *CardRepository) Init(){
	r.Kind = "Card"
	r.PutFn = func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
		card := po.(CardPO)
		return  datastore.Put(ctx, key, &card)
	}
	r.GetFn = func(ctx appengine.Context, keys []*datastore.Key) (ret []interface{}, err error) {
		var entities []CardPO
		entities = make([]CardPO, len(keys))
		err = datastore.GetMulti(ctx, keys, entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
	r.GetAllFn = func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
		var entities []CardPO
		keys, err = q.GetAll(ctx, &entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
}



type CardSuitRepository struct {
    tool.SimpleDAO
}

func (r *CardSuitRepository) Init(){
	r.Kind = "CardSuit"
	r.PutFn = func(ctx appengine.Context, key *datastore.Key, po interface{}) (retkey *datastore.Key, err error) {
		card := po.(CardSuitPO)
		return  datastore.Put(ctx, key, &card)
	}
	r.GetFn = func(ctx appengine.Context, keys []*datastore.Key) (ret []interface{}, err error) {
		var entities []CardSuitPO
		entities = make([]CardSuitPO, len(keys))
		err = datastore.GetMulti(ctx, keys, entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
	r.GetAllFn = func(ctx appengine.Context, q *datastore.Query) (ret []interface{}, keys []*datastore.Key, err error ) {
		var entities []CardSuitPO
		keys, err = q.GetAll(ctx, &entities)
		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
}