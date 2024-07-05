package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	fmt.Println(factorial(number))
}
