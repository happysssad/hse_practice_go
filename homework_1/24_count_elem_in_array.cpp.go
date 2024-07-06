func countElem(elem int, array []int) bool {
	cou := 0
	for _, value := range array {
		if elem == value {
		cou++
		}
	}
	return cou
}
