package pool

import (
	"os"
	"sync"
	"testing"
)

func BenchmarkBytesLocalPool(b *testing.B) {
	b.N = 16 * (b.N/16 + 1)
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(16)
	for p := 0; p < 16; p++ {
		go func(p int) {
			for i := 0; i < b.N/16; i++ {
				buf := GetBytes(16<<p)
				_, err = f.Write(buf)
				if err != nil {
					panic(err)
				}
				PutBytes(buf)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
func BenchmarkMake(b *testing.B) {
	b.N = 16 * (b.N/16 + 1)
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(16)
	for i := 0; i < 16; i++ {
		go func(p int) {
			for i := 0; i < b.N/16; i++ {
				buf := make([]byte, 16 << p)
				_, err = f.Write(buf)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
