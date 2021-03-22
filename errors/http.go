package errors

type HttpError struct {
	Msg  string
	Stat int
}

func (h *HttpError) Error() string {
	return h.Msg
}

func NewHttpError(msg string, stat int) *HttpError {
	return &HttpError{
		Msg:  msg,
		Stat: stat,
	}
}
