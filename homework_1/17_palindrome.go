package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	fmt.Println(isPalindrome(s))
}
