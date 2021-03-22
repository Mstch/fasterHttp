package pool

import (
	"math/rand"
	"sync"
	"testing"
)


func BenchmarkBytesLocalPool(b *testing.B) {
	b.N = 16 * (b.N/16 + 1)
	wg := &sync.WaitGroup{}
	wg.Add(16)
	for p := 0; p < 16; p++ {
		go func(p int) {
			s := rand.NewSource(int64(p))
			r := rand.New(s)
			for i := 0; i < b.N/16; i++ {
				buf := GetBytes(16 << r.Int31n(16))
				PutBytes(buf)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()
}
func BenchmarkSyncPool(b *testing.B) {
	b.N = 16 * (b.N/16 + 1)
	pools := make([]*sync.Pool, 16)
	for i := range pools {
		pools[i] = &sync.Pool{New: func() interface{} {
			return make([]byte, 16<<i)
		}}
	}
	wg := &sync.WaitGroup{}
	wg.Add(16)
	for i := 0; i < 16; i++ {
		go func(p int) {
			s := rand.NewSource(int64(p))
			r := rand.New(s)
			for i := 0; i < b.N/16; i++ {
				pi := r.Int31n(16)
				buf := pools[pi].Get().([]byte)
				pools[pi].Put(buf)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
