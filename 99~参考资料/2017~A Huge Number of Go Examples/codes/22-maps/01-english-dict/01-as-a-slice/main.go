package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	english := []string{"good", "great", "perfect"}
	turkish := []string{"iyi", "harika", "mükemmel"}

	// O(n) -> Inefficient: Depends on 'n'.
	for i, w := range english {
		if query == w {
			fmt.Printf("%q means %q\n", w, turkish[i])
			return
		}
	}

	fmt.Printf("%q not found\n", query)
}
