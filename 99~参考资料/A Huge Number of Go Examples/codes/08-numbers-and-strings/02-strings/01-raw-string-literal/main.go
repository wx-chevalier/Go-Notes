
package main

import "fmt"

func main() {
	// The type of a string and a raw string literal
	// is the same. They both are strings.
	//
	// So, they both can be used as a string value.
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	// string literal
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	// raw string literal
	s = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(s)

	// windows path
	fmt.Println("c:\\my\\dir\\file") // string literal
	fmt.Println(`c:\my\dir\file`)    // raw string literal
}
