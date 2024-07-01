package piscine

func IsNumeric(s string) bool {
	srune := []rune(s)
	i := 0
	for range srune {
		if srune[i] <= '9' && srune[i] >= '0' {
			i++
		} else {
			return false
		}
	}
	return true
}
