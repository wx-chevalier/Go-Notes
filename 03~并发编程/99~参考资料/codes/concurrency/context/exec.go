package context

import (
	"golang.org/x/net/context"
	"fmt"
	"math/rand"
	"time"
)

func Exec() {
	ctx := context.Background()
	ctx = Setup(ctx)

	rand.Seed(time.Now().UnixNano())

	timeoutCtx, cancel := context.WithTimeout(ctx, (time.Duration(rand.Intn(2)) * time.Millisecond))
	defer cancel()

	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(rand.Intn(2))*time.Millisecond))
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println(GetValue(ctx, timeoutKey))
			return
		case <-deadlineCtx.Done():
			fmt.Println(GetValue(ctx, deadlineKey))
			return
		}
	}
}
