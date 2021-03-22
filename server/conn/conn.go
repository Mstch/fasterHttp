package conn

import (
	"io"
	"net"
	"slowhttp/buffer"
	"slowhttp/constant"
	"slowhttp/mux"
	"slowhttp/pool"
	"slowhttp/req"
	"slowhttp/resp"
	"slowhttp/utils"
	"strconv"
)

type (
	HttpConn struct {
		net.Conn
		ReadBuf     *buffer.ReadBuffer
		off         int
		m           *mux.Mux
		writeWaiter *utils.Waiter
	}
)

func CreateHttpConn(c net.Conn, m *mux.Mux) *HttpConn {
	return &HttpConn{
		Conn:        c,
		ReadBuf:     &buffer.ReadBuffer{Buffer: *buffer.NewBuffer(make([]byte, 0, 512))},
		m:           m,
		writeWaiter: utils.NewWaiter(),
	}
}

func (hc *HttpConn) Serve() {
	go hc.WriteLoop()
	go hc.ReadLoop()
}
func (hc *HttpConn) ReadLoop() {
	reqId := uint64(0)
	for {
		curReq := &req.Request{
			Headers: make(map[string]string),
		}
		err := hc.ReadBuf.ReadAReqFrom(curReq, hc.Conn)
		if err != nil {
			hc.Close()
			if err != io.EOF {
				//panic(err)
			}
			return
		}
		go hc.handle(curReq, reqId)
		reqId++
	}
}
func (hc *HttpConn) WriteLoop() {
	respId := uint64(0)
	for {
		_, err := hc.Write(hc.writeWaiter.WaitFor(respId).([]byte))
		if err != nil {
			hc.Close()
			panic(err)
			return
		}
		respId++
	}
}

func (hc *HttpConn) handle(req *req.Request, reqId uint64) {
	defer release(req)
	//res := pool.RespPool.Get().(*resp.Response)
	res := &resp.Response{
		Version: constant.HttpVersion11,
		Headers: make(map[string]string),
		Body:    make([]byte, 0),
	}
	handle := hc.m.Match(req)
	if handle == nil{
		return
	}

	//this is concurrency
	handle(req, res)
	if res.Status == 0 {
		res.Status = 200
	}
	res.ContentLength = len(res.Body)

	//this is serial
	//<-pre
	err := hc.WriteResponse(res, reqId)
	//pool.RespPool.Put(res)
	pool.PutBytes(res.Body)
	if err != nil {
		panic(err)
	}
}

func release(req *req.Request) {
	pool.PutBytes(utils.ConvStringToBytes(req.Path))
	for _, v := range req.Headers {
		pool.PutBytes(utils.ConvStringToBytes(v))
	}
	//pool.ReqPool.Put(req)
}

func (hc *HttpConn) WriteResponse(res *resp.Response, id uint64) error {
	res.ContentLength = len(res.Body)
	res.Headers[constant.ContentLength] = strconv.Itoa(res.ContentLength)
	size := 0
	if statLine, ok := constant.RespStatusLine[res.Status]; ok {
		size += len(statLine) + 2
	} else {
		//todo
	}
	for k, v := range res.Headers {
		size += len(k) + 2 + len(v) + 2
	}
	size += 2
	size += len(res.Body)
	buf := pool.GetBytes(size)
	i := 0
	if statLine, ok := constant.RespStatusLine[res.Status]; ok {
		copy(buf, statLine)
		i += len(statLine)
		copy(buf[i:], constant.CRLF)
		i += 2
	} else {
		//todo
	}
	for k, v := range res.Headers {
		copy(buf[i:], utils.ConvStringToBytes(k))
		i += len(k)
		copy(buf[i:], constant.ColonSpace)
		i += 2
		copy(buf[i:], utils.ConvStringToBytes(v))
		i += len(v)
		copy(buf[i:], constant.CRLF)
		i += 2
	}
	copy(buf[i:], constant.CRLF)
	i += 2
	copy(buf[i:], res.Body)
	hc.writeWaiter.Join(id, buf)
	return nil
}
