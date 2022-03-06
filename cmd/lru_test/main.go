package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/dagrons/gin-demo/container/lru_cap"
)

func main() { // 压力测试
	cpufile, err := os.Create("cpu.prof")
	if err != nil {
		return
	}
	memfile, err := os.Create("mem.prof")
	if err != nil {
		return
	}
	pprof.StartCPUProfile(cpufile)
	pprof.WriteHeapProfile(memfile)
	l := lru_cap.New(1000) // cap = 1000
	for i := 0; i < 100; i++ {
		go func() { // random write
			for { // 10000内读写
				x := rand.Intn(10000)
				if rand.Intn(2) == 1 {
					l.Put(strconv.Itoa(x), x)
				} else {
					l.Put(strconv.Itoa(x), x+1)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}()

		go func() { // random read
			for {
				x := rand.Intn(10000)
				l.Get(strconv.Itoa(x))
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}
	time.Sleep(20 * time.Second)
	pprof.StopCPUProfile()
}
