package globalhelpers

func Contains(s []int, e int) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
