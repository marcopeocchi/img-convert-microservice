package internal

type Request struct {
	fn func() []byte
	c  chan []byte
}

func NewRequest(resultChan chan []byte, fn func() []byte) Request {
	return Request{
		fn: fn,
		c:  resultChan,
	}
}
