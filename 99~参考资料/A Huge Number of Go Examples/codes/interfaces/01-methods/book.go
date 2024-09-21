
package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	// b is a copy of the original `book` value here.
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}

// ----------------------------------------------------------------------------
// + you can use the same method names among different types.
// + you don't need to type `printGame`, it's just: `print`.
//
// func (b book) printBook() {
// 	// b is a copy of the original `book` value here.
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }

// ----------------------------------------------------------------------------
// b is a copy of the original `book` value here.
//
// func printBook(b book) {
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }
