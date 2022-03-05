package lrf

import (
	"container/list"
)

type Pair struct {
	K, V, F int
}

type LFUCache struct {
	Map     map[int]*list.Element
	ListMap map[int]*list.List
	Cap     int
	Min     int // min frequency
}

func (l *LFUCache) Get(k int) int {
	if e, ok := l.Map[k]; ok {
		return e.Value.(Pair).V
	}
	return -1
}

func (l *LFUCache) Put(key, val int) {
	if e, ok := l.Map[key]; ok {
		v := e.Value.(Pair)
		l.ListMap[v.F].Remove(e)
		if l.ListMap[v.F].Len() == 0 {
			delete(l.ListMap, v.F)
			if l.Min == v.F {
				l.Min++
			}
		}
		v.V = val
		v.F++
		el, ok := l.ListMap[v.F]
		if !ok {
			l.ListMap[v.F] = list.New()
		}
		el.PushFront(v)
	} else {
		el, ok := l.ListMap[1]
		if !ok {
			l.ListMap[1] = list.New()
		}
		v := Pair{key, val, 1}
		e := el.PushBack(v)
		l.Map[key] = e
		l.Min = 1 // 新增元素频率为1
	}
	if len(l.Map) > l.Cap { // 一定是新增元素造成的，所以Min一定为1
		el := l.ListMap[l.Min]
		e := el.Front()
		delete(l.Map, e.Value.(Pair).K)
		el.Remove(e)
		if el.Len() == 0 {
			delete(l.ListMap, l.Min)
		}
		l.Min = 1
	}
}
