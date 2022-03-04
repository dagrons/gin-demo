package lru

import "container/list"

type Pair struct {
	K, V int
}

type LRUCache struct {
	Map  map[int]*list.Element
	List *list.List
	Cap  int
}

func Constructor(cap int) LRUCache {
	return LRUCache{map[int]*list.Element{}, list.New(), cap}

}

func (l *LRUCache) Get(k int) int {
	if e, ok := l.Map[k]; ok {
		l.List.MoveToFront(e)
		return e.Value.(Pair).V
	}
	return -1
}

func (l *LRUCache) Put(k, v int) {
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
