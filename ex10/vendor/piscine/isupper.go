package piscine

func IsUpper(s string) bool {
	srune := []rune(s)
	i := 0
	for range srune {
		if srune[i] <= 'Z' && srune[i] >= 'A' {
			i++
		} else {
			return false
		}
	}
	return true
}