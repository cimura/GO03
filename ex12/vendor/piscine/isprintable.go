package piscine

func IsPrintable(s string) bool {
	srune := []rune(s)
	i := 0
	for range srune {
		if srune[i] <= '~' && srune[i] >= '!' {
			i++
		} else {
			return false
		}
	}
	return true
}
