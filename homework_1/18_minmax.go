func avarage(array []int) (int, int) {
	var min = array[0]
	var max = array[0]
	for _, value := range array {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	return min, max
}