package lru_cap

import (
	"testing"
	"time"
)

func TestLRU(t *testing.T) {
	lru := New(2)
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
