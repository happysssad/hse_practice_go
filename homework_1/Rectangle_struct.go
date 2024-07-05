package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var width, height float64
	_, err := fmt.Scanf("%f %f", &width, &height)
	if err != nil {
		return
	}
	var x = Rectangle{width: width, height: height}
	var area = x.Area()
	fmt.Println(area)
}
