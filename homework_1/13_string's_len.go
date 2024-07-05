package main

import "fmt"

func getLength(s string) int {
	var count = 0
	for range s {
		count++
	}
	return count
}

func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	fmt.Println(getLength(s))
}
