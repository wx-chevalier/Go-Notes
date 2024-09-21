package reactive

import (
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/subscription"
)

// Exec connects rxgo and returns
// our results side-effect + a subscription
// channel to block on at the end
func Exec() (Results, <-chan subscription.Subscription) {
	results := make(Results)
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			wine, ok := item.(Wine)
			if ok {
				result := results[wine.Age]
				result.SumRating += wine.Rating
				result.NumSamples++
				results[wine.Age] = result
			}
		},
	}
	wine := GetWine()
	it, _ := iterable.New(wine)

	source := observable.From(it)
	sub := source.Subscribe(watcher)

	return results, sub
}
