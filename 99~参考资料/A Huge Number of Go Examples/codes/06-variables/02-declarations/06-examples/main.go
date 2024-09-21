
package main

import "fmt"

func main() {
	// names are case-sensitive:
	// MyAge, myAge, and MYAGE are different variables

	// USE-CASE:
	// When to use a parallel declaration?
	//
	// NOT GOOD:
	// var myAge int
	// var yourAge int
	//
	// SO-SO:
	// var (
	// 	myAge int
	// 	yourAge int
	// )
	//
	// BETTER:
	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var temperature float64
	fmt.Println(temperature)

	var success bool
	fmt.Println(success)

	var language string
	fmt.Println(language)
}
