
package main

func main() {
	var speed int
	speed = 50 // OK

	var running bool
	running = true // OK

	var force float64
	speed = int(force)

	_, _, _ = speed, running, force
}
