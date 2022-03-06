package lru_cap

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestLRU(t *testing.T) { // 正确性测试
	lru := New(2, 256, 256)
	lru.Put("1", 2)
	time.Sleep(100 * time.Millisecond)
	if lru.Get("2") != -1 {
		t.Fail()
	}
	if lru.Get("1") != 2 {
		t.Fail()
	}
	lru.Put("2", 3)
	time.Sleep(100 * time.Millisecond)
	if lru.Get("2") != 3 {
		t.Fail()
	}
	lru.Put("3", 4)
	time.Sleep(100 * time.Millisecond)
	if lru.Get("3") != 4 {
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	if lru.Get("1") != -1 {
		t.Fail()
	}
}

func BenchmarkLRU(b *testing.B) { // 压力测试
	l := New(1000, 256, 256) // cap = 1000
	for i := 0; i < 100000; i++ {
		x := rand.Intn(100000)
		if rand.Intn(2) == 1 {
			l.Put(strconv.Itoa(x), x)
		} else {
			l.Put(strconv.Itoa(x), x+1)
		}
		x = rand.Intn(100000)
		l.Get(strconv.Itoa(x))
	}
}
