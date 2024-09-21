
package main

import "fmt"

func main() {
	var counter int

	counter++
	counter--

	counter += 5
	counter *= 10
	counter /= 5

	fmt.Println(counter)
}
