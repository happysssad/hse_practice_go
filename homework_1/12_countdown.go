package main

import "fmt"

func main() {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	for i := 0; i < number; i++ {
		fmt.Printf("%d ", number-i)
	}
}
