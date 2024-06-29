package piscine

func Index(s string, toFind string) int {
	for i := range s {
		if rune(s[i]) == rune(toFind[0]) {
			return i
		}
	}
	return -1
}