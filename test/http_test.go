package test

import (
	"net"
	"net/http"
	_ "net/http/pprof"
	"slowhttp/pool"
	"slowhttp/req"
	"slowhttp/resp"
	"slowhttp/server"

	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkSlowHTTP(b *testing.B) {
	s, err := server.NewServer(":8888")
	if err != nil {
		panic(err)
	}
	defer s.Shutdown()
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	total := uint64(0)
	resetTimerO := &sync.Once{}
	s.RegHandler("/123456", func(req *req.Request, res *resp.Response) {
		if atomic.AddUint64(&total, 1) > uint64(b.N) {
			println(total, b.N)
			return
		}
		res.Body = pool.GetBytes(len(req.Path))
		copy([]byte(req.Path), res.Body)
		resetTimerO.Do(func() {
			b.ResetTimer()
		})
		wg.Done()
	})
	go s.Serve()
	wg.Wait()
}

func BenchmarkStdHTTP(b *testing.B) {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	s := http.NewServeMux()
	s.HandleFunc("/123456", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RequestURI))
	})
	go http.Serve(l, s)
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://127.0.0.1:8888/123456")
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}

func TestSlowHTTP(t *testing.T) {
	s, err := server.NewServer(":8888")
	if err != nil {
		panic(err)
	}
	defer s.Shutdown()
	s.RegHandler("/123456", func(req *req.Request, res *resp.Response) {
		res.Headers["Content-Type"] = "text/plain"
		res.Body = pool.GetBytes(len(req.Path))
		copy(res.Body, req.Path)
	})
	go s.Serve()
	http.ListenAndServe(":9999", nil)
}
func TestStdHTTP(t *testing.T) {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	s := http.NewServeMux()
	s.HandleFunc("/123456", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RequestURI))
	})
	 http.Serve(l, s)
}
