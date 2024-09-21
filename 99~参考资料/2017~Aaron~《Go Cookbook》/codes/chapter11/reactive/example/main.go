package main

import (
	"fmt"

	"github.com/agtorre/go-cookbook/chapter11/reactive"
)

func main() {
	results, sub := reactive.Exec()

	// wait for the channel to emit a Subscription
	<-sub

	// process results
	for key, val := range results {
		fmt.Printf("Age: %d, Sample Size: %d, Average Rating: %.2f\n", key, val.NumSamples, val.SumRating/float64(val.NumSamples))
	}
}
