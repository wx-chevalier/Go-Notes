
package main

func main() {
	var speed int
	// speed = "100" // NOT OK

	var running bool
	// running = 50 // NOT OK

	var force float64
	// speed = force // NOT OK

	_, _, _ = speed, running, force
}
