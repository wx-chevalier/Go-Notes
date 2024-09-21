package pipeline

import (
	"golang.org/x/net/context"
	"encoding/base64"
	"fmt"
)

func (w *Worker) Encode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			w.out <- fmt.Sprintf("%s => %s", val, base64.StdEncoding.EncodeToString([]byte(val)))
		}
	}
}
