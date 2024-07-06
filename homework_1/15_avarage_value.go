func avarage(array []int) int {
	var sum = 0
	for _, value := range array {
		sum += value
	}
	return float64(sum) / float64(len(array))
}
