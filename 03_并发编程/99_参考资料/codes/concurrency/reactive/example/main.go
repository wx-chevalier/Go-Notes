package main

import (
	"fmt"
	"reactive"
)

func main() {
	results, sub := reactive.Exec()

	<-sub

	for key, val := range results {
		fmt.Printf("Age: %d, Sample Size: %d, Average Rating: %.2f\n", key, val.NumSamples, val.SumRating/float64(val.NumSamples))
	}
}
