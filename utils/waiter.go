package utils

import "sync"

type (
	Waiter struct {
		chPool *sync.Pool
		wMap   *sync.Map
	}
)

func NewWaiter() *Waiter {
	return &Waiter{
		chPool: &sync.Pool{New: func() interface{} {
			return make(chan interface{}, 1)
		}},
		wMap: &sync.Map{},
	}
}
func (w *Waiter) Join(n uint64, v interface{}) {
	ch, _ := w.wMap.LoadOrStore(n, w.chPool.Get().(chan interface{}))
	ch.(chan interface{}) <- v
}
func (w *Waiter) WaitFor(n uint64) interface{} {
	ch, _ := w.wMap.LoadOrStore(n, w.chPool.Get().(chan interface{}))
	v := <-ch.(chan interface{})
	w.chPool.Put(ch)
	return v
}
