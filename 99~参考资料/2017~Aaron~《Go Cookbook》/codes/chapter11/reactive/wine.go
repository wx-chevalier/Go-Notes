package reactive

// Wine represents a bottle
// of wine and is our
// input stream
type Wine struct {
	Name   string
	Age    int
	Rating float64 // 1-5
}

// GetWine returns an array of wines,
// ages, and ratings
func GetWine() interface{} {
	// some example wines
	w := []interface{}{
		Wine{"Merlot", 2011, 3.0},
		Wine{"Cabernet", 2010, 3.0},
		Wine{"Chardonnay", 2010, 4.0},
		Wine{"Pinot Grigio", 2009, 4.5},
	}
	return w
}

// Results holds a list of results by age
type Results map[int]Result

// Result is used for aggregation
type Result struct {
	SumRating  float64
	NumSamples int
}
