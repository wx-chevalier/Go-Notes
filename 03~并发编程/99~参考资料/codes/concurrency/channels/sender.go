package channels

import "time"

func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-done:
			ch <- "sender done."
			return
		case <-t:
			ch <- "tick"
		}
	}
}
