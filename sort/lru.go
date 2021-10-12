package sort

import (
	"container/list"
	"sync"
)

type Lru struct {
	max      int
	dataList *list.List
	dataMap  map[interface{}]*list.Element
	rwLock   sync.RWMutex
}

type Node struct {
	Key interface{}
	Val interface{}
}

func NewLru(max int) *Lru {
	return &Lru{
		max:      max,
		dataList: list.New(),
		dataMap:  make(map[interface{}]*list.Element, 0),
	}
}

func (z *Lru) Add(key, val interface{}) {
	if e, ok := z.dataMap[key]; ok {
		e.Value.(*Node).Val = e
		z.dataList.MoveToFront(e)
		return
	}

	e := z.dataList.PushFront(&Node{Key: key, Val: val})
	z.dataMap[key] = e

	if z.dataList.Len() > z.max {
		e := z.dataList.Back()
		z.dataList.Remove(e)
		delete(z.dataMap, key)
	}
	return

}

func (z *Lru) Get(key interface{}) interface{} {

	if e, ok := z.dataMap[key]; ok {
		z.dataList.MoveToFront(e)
		return e.Value.(*Node).Val
	}

	return nil

}

func (z *Lru) Remove(key interface{}) {
	if e, ok := z.dataMap[key]; ok {
		delete(z.dataMap, key)
		z.dataList.Remove(e)
	}
}
