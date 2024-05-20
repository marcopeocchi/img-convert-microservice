package internal

type Worker struct {
	requests chan Request // conversions to do
	pending  int          // conversions pending
	index    int          // index in the heap
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}
