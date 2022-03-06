package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/dagrons/gin-demo/container/lru"
)

func main() { // 压力测试
	startTime := time.Now()
	sizeMove, sizeDel := 1, 1
	l := lru.New(10000) // cap = 1000
	cpufile, err := os.Create("cpu" + fmt.Sprintf("size_move_%d_size_del_%d", sizeMove, sizeDel) + ".prof")
	if err != nil {
		return
	}
	memfile, err := os.Create("mem.prof" + fmt.Sprintf("size_move_%d_size_del_%d", sizeMove, sizeDel) + ".prof")
	if err != nil {
		return
	}
	cnt := []int64{0}
	pprof.StartCPUProfile(cpufile)
	pprof.WriteHeapProfile(memfile)
	for i := 0; i < 1000000; i++ {
		go func() { // random write
			for { // 10000内读写
				x := rand.Intn(100000)
				if rand.Intn(2) == 1 {
					l.Put(strconv.Itoa(x), x)
				} else {
					l.Put(strconv.Itoa(x), x+1)
				}
				atomic.AddInt64(&cnt[0], 1)
				if cnt[0] > 10000000 {
					return
				}
			}
		}()

		go func() { // random read
			for {
				x := rand.Intn(100000)
				l.Get(strconv.Itoa(x))
				atomic.AddInt64(&cnt[0], 1)
				if cnt[0] > 10000000 {
					return
				}
			}
		}()
	}
	for {
		if cnt[0] > 10000000 {
			pprof.StopCPUProfile()
			fmt.Printf("cost time=%v\n", time.Since(startTime))
			return
		}
	}
}
