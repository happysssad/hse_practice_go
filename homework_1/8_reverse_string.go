func reverseString(s string) string {
	var ans string
	for _, v := range s {
		ans = string(v) + ans
	}
	return ans
}