package palindromestrings

import "github.com/Dev-Snippets/algorithm-go-snippets/numbers/reverses"

// IsPalindrome determines if the input is a palindrome
func IsPalindrome(number int) bool {
	return number == reverses.Reverse(number)
}
