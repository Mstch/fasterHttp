package utils

import (
	"runtime"
	"testing"
)

func TestProcPin(t *testing.T) {
	total := make([]int, runtime.GOMAXPROCS(0))
	done := make(chan struct{}, 256)
	for i := 0; i < 256; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				total[ProcPin()]++
				ProcUnpin()
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 256; i++ {
		<-done
	}
	sum := 0
	for _, s := range total {
		sum += s
	}
	if sum != 256*10000 {
		panic(sum)
	}
}
