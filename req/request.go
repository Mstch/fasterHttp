package req

import (
	"net/url"
	"slowhttp/session"
)

type Request struct {
	Method        string
	Path          string
	Version       string
	Headers       map[string]string
	ContentLength int
	Body          []byte
	URL           *url.URL
	Session       *session.Session
}

func (r *Request) GetQuery(k string) string {
	return r.URL.Query().Get(k)
}
