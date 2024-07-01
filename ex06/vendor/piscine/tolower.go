package piscine

func ToLower(s string) string {
	srune := [] rune(s)
	i := 0
	for range s {
		if srune[i] <= 'Z' && srune[i] >= 'A' {
			srune[i] += 32
		}
		i++
	}
	return string(srune)
}