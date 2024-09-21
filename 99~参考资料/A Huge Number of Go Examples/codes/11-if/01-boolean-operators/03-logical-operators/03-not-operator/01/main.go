
package main

import "fmt"

func main() {
	fmt.Println(
		"hi" == "hi" && 3 > 2,    //   true  && true  => true
		"hi" != "hi" || 3 > 2,    //   false || true  => true
		!("hi" != "hi" || 3 > 2), // !(false || true) => false
	)
}
