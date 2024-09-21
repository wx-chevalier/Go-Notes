package reactive

type Wine struct {
	Name   string
	Age    int
	Rating float64 
}

func GetWine() interface{} {
	w := []interface{}{
		Wine{"Merlot", 2011, 3.0},
		Wine{"Cabarnet", 2010, 3.0},
		Wine{"Chardanay", 2010, 4.0},
		Wine{"Pinot Grigio", 2009, 4.5},
	}
	return w
}

type Results map[int]Result

type Result struct {
	SumRating  float64
	NumSamples int
}
