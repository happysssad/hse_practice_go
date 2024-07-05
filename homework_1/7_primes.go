package main

import "fmt"

func isprime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	for i := 2; i <= number; i++ {
		if isprime(i) {
			fmt.Println(i)
		}
	}
}
