package piscine

func ToUpper(s string) string {
	srune := [] rune(s)
	i := 0
	for range s {
		if srune[i] <= 'z' && srune[i] >= 'a' {
			srune[i] -= 32
		}
		i++
	}
	return string(srune)
}