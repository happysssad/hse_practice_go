package main

import "fmt"

func fibbonachi(n int) []int {
	f := make([]int, n)
	f[0] = 0
	if n == 1 {
		return f
	}
	f[1] = 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
func main() {
	var number int
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	fmt.Println(fibbonachi(number))
}
