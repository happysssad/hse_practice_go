package main

import "fmt"

func even_or_odd(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func main() {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	if even_or_odd(number) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}
