package constant

const (
	MethodGET     = "GET"
	MethodHEAD    = "HEAD"
	MethodPOST    = "POST"
	MethodPUT     = "PUT"
	MethodDELETE  = "DELETE"
	MethodCONNECT = "CONNECT"
	MethodOPTIONS = "OPTIONS"
	MethodPATCH   = "PATCH"
)

const (
	HttpVersion11 = "HTTP/1.1"
)

const (
	MaxHeaderLineLen = 8192
)

var (
	CRLF = []byte{'\r', '\n'}

	ColonSpace = []byte{':', ' '}
)
var (
	ValidMethod = map[string]string{
		MethodGET:     MethodGET,
		MethodHEAD:    MethodHEAD,
		MethodPOST:    MethodPOST,
		MethodPUT:     MethodPUT,
		MethodDELETE:  MethodDELETE,
		MethodCONNECT: MethodCONNECT,
		MethodOPTIONS: MethodOPTIONS,
		MethodPATCH:   MethodPATCH,
	}
	RespStatusLine = map[int][]byte{
		100: []byte("HTTP/1.1 100 Continue"),
		101: []byte("HTTP/1.1 101 Switching Protocols"),
		200: []byte("HTTP/1.1 200 OK"),
		201: []byte("HTTP/1.1 201 Created"),
		202: []byte("HTTP/1.1 202 Accepted"),
		203: []byte("HTTP/1.1 203 Non-Authoritative Information"),
		204: []byte("HTTP/1.1 204 No Content"),
		205: []byte("HTTP/1.1 205 Reset Content"),
		206: []byte("HTTP/1.1 206 Partial Content"),
		300: []byte("HTTP/1.1 300 Multiple Choices"),
		301: []byte("HTTP/1.1 301 Moved Permanently"),
		302: []byte("HTTP/1.1 302 Found"),
		303: []byte("HTTP/1.1 303 See Other"),
		304: []byte("HTTP/1.1 304 Not Modified"),
		305: []byte("HTTP/1.1 305 Use Proxy"),
		307: []byte("HTTP/1.1 307 Temporary Redirect"),
		400: []byte("HTTP/1.1 400 Bad Request"),
		401: []byte("HTTP/1.1 401 Unauthorized"),
		402: []byte("HTTP/1.1 402 Payment Required"),
		403: []byte("HTTP/1.1 403 Forbidden"),
		404: []byte("HTTP/1.1 404 Not Found"),
		405: []byte("HTTP/1.1 405 Method Not Allowed"),
		406: []byte("HTTP/1.1 406 Not Acceptable"),
		407: []byte("HTTP/1.1 407 Proxy Authentication Required"),
		408: []byte("HTTP/1.1 408 Request Time-out"),
		409: []byte("HTTP/1.1 409 Conflict"),
		410: []byte("HTTP/1.1 410 Gone"),
		411: []byte("HTTP/1.1 411 Length Required"),
		412: []byte("HTTP/1.1 412 Precondition Failed"),
		413: []byte("HTTP/1.1 413 Request Entity Too Large"),
		414: []byte("HTTP/1.1 414 Request-URI Too Large"),
		415: []byte("HTTP/1.1 415 Unsupported Media Type"),
		416: []byte("HTTP/1.1 416 Requested range not satisfiable"),
		417: []byte("HTTP/1.1 417 Expectation Failed"),
		500: []byte("HTTP/1.1 500 Internal Server Error"),
		501: []byte("HTTP/1.1 501 Not Implemented"),
		502: []byte("HTTP/1.1 502 Bad Gateway"),
		503: []byte("HTTP/1.1 503 Service Unavailable"),
		504: []byte("HTTP/1.1 504 Gateway Time-out"),
		505: []byte("HTTP/1.1 505 HTTP Version not supported"),
	}
)
