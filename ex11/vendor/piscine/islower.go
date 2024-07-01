package piscine

func IsLower(s string) bool {
	srune := []rune(s)
	i := 0
	for range srune {
		if srune[i] <= 'z' && srune[i] >= 'a' {
			i++
		} else {
			return false
		}
	}
	return true
}