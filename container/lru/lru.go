package lru

import (
	"container/list"
	"sync"
)

type Pair struct {
	K string
	V interface{}
}
type LRUCache struct {
	Map  map[string]*list.Element
	List *list.List
	Cap  int
	Lock *sync.RWMutex
}

func New(cap int) *LRUCache {
	return &LRUCache{map[string]*list.Element{}, list.New(), cap, &sync.RWMutex{}}
}

func (l *LRUCache) Get(k string) interface{} {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	if e, ok := l.Map[k]; ok {
		l.List.MoveToFront(e)
		return e.Value.(Pair).V
	}
	return -1
}

func (l *LRUCache) Put(k string, v interface{}) {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	if e, ok := l.Map[k]; ok {
		e.Value = Pair{k, v}
		l.List.MoveToFront(e)
	} else {
		e := l.List.PushFront(Pair{k, v})
		l.Map[k] = e
	}
	if len(l.Map) > l.Cap {
		e := l.List.Back()
		l.List.Remove(e)
		delete(l.Map, e.Value.(Pair).K)
	}
}
