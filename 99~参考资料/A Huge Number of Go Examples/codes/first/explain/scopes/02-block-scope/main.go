
package main

func nope() { // block scope starts

	// hello and ok are only visible here
	const ok = true
	var hello = "Hello"

	_ = hello
} // block scope ends

func main() { // block scope starts

	// hello and ok are not visible here

	// ERROR:
	// fmt.Println(hello, ok)

} // block scope ends
