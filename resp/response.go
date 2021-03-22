package resp

type Response struct {
	Version       string
	Status        int
	Headers       map[string]string
	Body          []byte
	ContentLength int
}
