package pipeline

import "golang.org/x/net/context"

type Worker struct {
	in  chan string
	out chan string
}

type Job string

const (
	Print Job = "print"
	Encode Job = "encode"
)
func (w *Worker) Work(ctx context.Context, j Job) {
	switch j {
	case Print:
		w.Print(ctx)
	case Encode:
		w.Encode(ctx)
	default:
		return
	}
}
