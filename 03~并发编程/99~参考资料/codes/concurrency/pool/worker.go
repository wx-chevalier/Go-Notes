package pool

import (
	"golang.org/x/net/context"
	"fmt"
)

func Dispatch(numWorker int) (context.CancelFunc, chan WorkRequest, chan WorkResponse) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	in := make(chan WorkRequest, 10)
	out := make(chan WorkResponse, 10)

	for i := 0; i < numWorker; i++ {
		go Worker(ctx, i, in, out)
	}
	return cancel, in, out
}

func Worker(ctx context.Context, id int, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			fmt.Printf("worker id: %d, performing %s work\n", id, wr.Op)
			out <- Process(wr)
		}
	}
}
