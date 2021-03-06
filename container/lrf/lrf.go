package lrf

import "container/list"

type PIII struct {
	K, V, F int
}

type LFUCache struct {
	Map     map[int]*list.Element
	ListMap map[int]*list.List
	Cap     int
	Min     int
}

func Constructor(cap int) LFUCache {
	return LFUCache{map[int]*list.Element{}, map[int]*list.List{}, cap, 0}
}

func (l *LFUCache) Get(k int) int {
	e, ok := l.Map[k]
	if !ok { // 不存在，直接返回-1
		return -1
	}
	// 存在，上升到上一级列表
	v := e.Value.(*PIII)
	l.ListMap[v.F].Remove(e)
	v.F++
	if _, ok := l.ListMap[v.F]; !ok {
		l.ListMap[v.F] = list.New()
	}
	l.Map[k] = l.ListMap[v.F].PushBack(v)
	if l.Min == v.F-1 && l.ListMap[v.F-1].Len() == 0 {
		l.Min++
	}
	return v.V
}

func (l *LFUCache) Put(key, val int) {
	if l.Cap == 0 {
		return
	}
	if e, ok := l.Map[key]; ok { // 已经存在，更新即可
		v := e.Value.(*PIII)
		v.V = val
		l.Get(key)
		return
	} // 还不存在，首先判断容量，再创建并加入
	if len(l.Map) == l.Cap {
		e := l.ListMap[l.Min].Front()
		delete(l.Map, e.Value.(*PIII).K)
		l.ListMap[l.Min].Remove(e)
	}
	l.Min = 1
	nv := &PIII{key, val, 1} // new value
	if _, ok := l.ListMap[1]; !ok {
		l.ListMap[1] = list.New()
	}
	l.Map[key] = l.ListMap[1].PushBack(nv)
}

// 什么时候会加key，什么时候会减key
