func linSearch(elem int, array []int) int {
	for ind, value := range array {
		if elem == value {
			return ind
		}
	}
	return -1
}
