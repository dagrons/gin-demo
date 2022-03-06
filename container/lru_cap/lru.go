// http://10.112.108.112:3000/HblS-xIuQmaO2YJFH8Pwgg

package lru_cap

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync"
)

type Pair struct {
	K   string
	V   interface{}
	cmd int
}

const (
	MovePair = iota
	AddPair
)

type LRUBucket struct {
	Lock *sync.RWMutex
	Map  map[string]*list.Element
}

type LRUCache struct {
	Buckets     []*LRUBucket // 哈希分桶
	List        *list.List
	Cap         int
	movePairs   chan *list.Element
	deletePairs chan struct{}
}

func (l *LRUCache) Len() int {
	totalNum := 0
	for _, bucket := range l.Buckets {
		totalNum += len(bucket.Map)
	}
	return totalNum
}

func New(cap int) *LRUCache {
	l := &LRUCache{}
	l.Buckets = make([]*LRUBucket, 256)
	for i := 0; i < 256; i++ {
		l.Buckets[i] = &LRUBucket{
			&sync.RWMutex{},
			map[string]*list.Element{},
		}
	}
	l.List = list.New()
	l.Cap = cap
	l.movePairs = make(chan *list.Element, 10)
	l.deletePairs = make(chan struct{}, 10)
	go l.worker() // 将链表相关操作解耦，独立线程处理
	return l
}

func getBucketIdx(key string) int {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	res, err := strconv.ParseInt(fmt.Sprintf("%x", hasher.Sum(nil))[0:2], 16, 0)
	if err != nil {
		panic(fmt.Errorf("error parseInt for %s", key))
	}
	return int(res)
}

func (l *LRUCache) Get(key string) interface{} {
	shardKey := getBucketIdx(key)
	l.Buckets[shardKey].Lock.RLock() // 读写临界区
	e, ok := l.Buckets[shardKey].Map[key]
	l.Buckets[shardKey].Lock.RUnlock()
	if !ok {
		return -1
	}
	l.movePairs <- e
	return e.Value.(Pair).V
}

func (l *LRUCache) Put(key string, val interface{}) {
	shardKey := getBucketIdx(key)
	l.Buckets[shardKey].Lock.RLock() // 读写临界区
	e, ok := l.Buckets[shardKey].Map[key]
	l.Buckets[shardKey].Lock.RUnlock()
	if ok {
		t := e.Value.(Pair)
		t.V = val
		l.movePairs <- e
	} else {
		newNode := &list.Element{Value: Pair{key, val, AddPair}}
		l.movePairs <- newNode
	}
}

func (l *LRUCache) worker() { // 事件驱动
	for {
		select {
		case e := <-l.movePairs:
			switch e.Value.(Pair).cmd {
			case AddPair:
				nVal := e.Value.(Pair)
				nVal.cmd = MovePair
				newNode := l.List.PushFront(nVal)
				nShardKey := getBucketIdx(nVal.K)
				l.Buckets[nShardKey].Lock.Lock()
				l.Buckets[nShardKey].Map[nVal.K] = newNode
				l.Buckets[nShardKey].Lock.Unlock()
				if l.Len() > l.Cap {
					node := l.List.Back()
					k := node.Value.(Pair).K
					shardKeyBackNode := getBucketIdx(k)
					l.Buckets[shardKeyBackNode].Lock.Lock()
					delete(l.Buckets[shardKeyBackNode].Map, k)
					l.Buckets[shardKeyBackNode].Lock.Unlock()
					select {
					case l.deletePairs <- struct{}{}:
					default:
						<-l.deletePairs
						l.List.Remove(l.List.Back())
					}
				}
			case MovePair:
				l.List.MoveToFront(e)
			}
		case <-l.deletePairs:
			l.List.Remove(l.List.Back())
		}
	}
}
