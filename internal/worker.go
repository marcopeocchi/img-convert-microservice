package internal

type Worker struct {
	requests chan Request // downloads to do
	pending  int          // downloads pending
	index    int          // index in the heap
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}
