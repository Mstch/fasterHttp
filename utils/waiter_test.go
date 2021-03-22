package utils

import "testing"

func BenchmarkWaiter(b *testing.B) {
	b.N = 32 * (b.N/32 + 1)
	w := NewWaiter()
	for i := 0; i < 32; i++ {
		go func() {
			for i := 0; i < b.N/32; i++ {
				w.Join(uint64(100), nil)
			}
		}()
	}
	for i := b.N - 1; i >= 0; i-- {
		w.WaitFor(uint64(100))
	}
}
