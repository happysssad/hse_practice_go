func removeElem(slice []int, ind int) []int {
	if ind >= len(slice) || ind < 0 {
		return slice
	}
	return append(slice[:ind], slice[ind+1:]...)
}
