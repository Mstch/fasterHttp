package pool

import (
	"slowhttp/constant"
	"slowhttp/req"
	"slowhttp/resp"
	"sync"
)

var (
	ReqPool = &sync.Pool{New: func() interface{} {
		return &req.Request{
			Headers:make(map[string]string),
		}
	}}
	RespPool = &sync.Pool{New: func() interface{} {
		return &resp.Response{
			Version: constant.HttpVersion11,
			Headers: make(map[string]string),
			Body:    make([]byte, 0),
		}
	}}
	ChanPool = &sync.Pool{New: func() interface{} {
		return make(chan struct{}, 1)
	}}
)
