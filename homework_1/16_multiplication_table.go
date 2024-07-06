package main

import "fmt"

func main() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	var n int
	for _, n = range numbers {
		fmt.Println(number, "x", n, "=", n*number)
	}
}
