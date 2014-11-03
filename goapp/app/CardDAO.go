package app

import (
	"lib/tool"
	"appengine"
    "appengine/datastore"
	"sort"
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
		var entities []CardEntity
		keys, err = q.GetAll(ctx, &entities)
		
		w := func(c1, c2 *CardEntity) bool {
			return c1.Weight < c2.Weight
		}
		weaponType := func(c1, c2 *CardEntity) bool {
			return c1.Type < c2.Type
		}
		level := func(c1, c2 *CardEntity) bool {
			return c1.Level < c2.Level
		}
		name := func(c1, c2 *CardEntity) bool {
			return c1.Name < c2.Name
		}
		atkType := func(c1, c2 *CardEntity) bool {
			return c1.AtkType < c2.AtkType
		}
		OrderedBy(w, weaponType, level, name, atkType).Sort(entities)

		for idx, entity := range entities {
			entity.Key = keys[idx].IntID()
			ret = append(ret, entity)
		}
		return
	}
}




type lessFunc func(p1, p2 *CardEntity) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	changes []CardEntity
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes []CardEntity) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}