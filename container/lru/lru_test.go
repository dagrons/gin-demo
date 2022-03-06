package lru

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkLRU(b *testing.B) { // 压力测试
	l := New(1000) // cap = 1000
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
