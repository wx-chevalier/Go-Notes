package main

import (
	"golang.org/x/net/context"
	"time"

	"github.com/agtorre/go-solutions/section2/channels"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go channels.Printer(ctx, ch)
	go channels.Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()
	time.Sleep(1 * time.Second)
}
