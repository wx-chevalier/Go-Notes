
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 🛑 DON'T DO THIS:
//     It's hard to read.
//     It's just an exercise.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}
