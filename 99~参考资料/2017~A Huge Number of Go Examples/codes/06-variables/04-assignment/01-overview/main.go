
package main

import "fmt"

func main() {
	var counter int

	fmt.Println("counter's name : counter")
	fmt.Println("counter's value:", counter)
	fmt.Printf("counter's type : %T\n", counter)

	counter = 10 // OK
	// counter = "ten" // NOT OK

	fmt.Println("counter's value:", counter)

	counter++
	fmt.Println("counter's value:", counter)
}
