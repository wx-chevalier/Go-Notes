package channels

import (
	"golang.org/x/net/context"
	"fmt"
	"time"
)

func Printer(ctx context.Context, ch chan string) {
	t := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printer done.")
			return
		case res := <-ch:
			fmt.Println(res)
		case <-t:
			fmt.Println("tock")
		}
	}
}

select{
 case <-time.Tick(200 * time.Millisecond):
}
