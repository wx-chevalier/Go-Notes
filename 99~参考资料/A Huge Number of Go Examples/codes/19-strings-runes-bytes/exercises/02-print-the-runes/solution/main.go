
package main

import "fmt"

func main() {
	const word = "console"

	for _, w := range word {
		fmt.Printf("%c\n", w)
		fmt.Printf("\tdecimal: %[1]d\n", w)
		fmt.Printf("\thex    : 0x%[1]x\n", w)
		fmt.Printf("\tbinary : 0b%08[1]b\n", w)
	}

	// print the word manually using runes
	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
