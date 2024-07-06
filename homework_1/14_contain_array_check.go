func containElemCheck(elem int, array []int) bool {
	for _, value := range array {
		if elem == value {
			return true
		}
	}
	return false
}

