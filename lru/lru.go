package lru

import (
	"container/list"
	"sync"
)

type Lru struct {
	dataList *list.List
	dataMap  map[interface{}]*list.Element
	max      int
	lock     sync.Mutex
}

type Node struct {
	key   interface{}
	value interface{}
}

func NewLru(max int) *Lru {
	return &Lru{
		dataList: &list.List{},
		dataMap:  make(map[interface{}]*list.Element),
		max:      max,
	}
}

func (z *Lru) Add(key, value interface{}) {
	if e, ok := z.dataMap[key]; ok {
		e.Value.(*Node).value = value
		z.dataList.MoveToFront(e)
	}

	e := z.dataList.PushFront(&Node{key: key, value: value})
	z.dataMap[key] = e
	if z.dataList.Len() > z.max {
		e := z.dataList.Back()
		delete(z.dataMap, e.Value.(*Node).key)
		z.dataList.Remove(e)
	}
}

func (z *Lru) Get(key interface{}) interface{} {
	if e, ok := z.dataMap[key]; ok {
		z.dataList.MoveToFront(e)
		return e.Value.(*Node).value
	}
	return nil
}
