
package main

import "fmt"

// This program uses magic values

func main() {
	cm := 100
	m := cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)
}
