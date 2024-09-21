package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
