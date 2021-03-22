package errors

var (
	ErrInvalidRequestLine   = NewHttpError("invalid request line", 400)
	ErrInvalidMethod        = NewHttpError("invalid method", 400)
	ErrInvalidPath          = NewHttpError("invalid path", 400)
	ErrInvalidHeader        = NewHttpError("invalid header", 400)
	ErrNoContentLength      = NewHttpError("no content length", 400)
	ErrInvalidContentLength = NewHttpError("invalid content length", 400)
	ErrInvalidVersion       = NewHttpError("invalid http version", 400)
	ErrConnError            = NewHttpError("conn err", -1)
)
