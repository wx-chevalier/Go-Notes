package main

import (
	"fmt"
)

var limit = make(chan int, 3)

type IWorker interface {
	test()
}

type Worker struct {
}

func (w *Worker) test() {
	fmt.Println("work")
}

func main_2() {

	work := [...](IWorker){&Worker{}, &Worker{}}

	for _, w := range work {
		go func() {
			limit <- 1
			w.test()
			<-limit
		}()
	}

	select {}
}
