
package main

import "fmt"

// You can't use declaration statements without a keyword
// Short declaration doesn't have a keyword (`var`)
// So, it can't be used at the package scope
//
// SYNTAX ERROR:
// "non-declaration statement outside function body"

// safe := true

// However, you can use the normal declaration at the
// package scope. Since it has a keyword: `var`
var safe = true

func main() {
	fmt.Println(safe)
}
