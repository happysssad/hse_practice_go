package main

import "fmt"

func CelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func main() {
	var t float64
	_, err := fmt.Scanf("%f", &t)
	if err != nil {
		return
	}
	fmt.Println(CelsiusToFahrenheit(t))
}
