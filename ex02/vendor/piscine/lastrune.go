package	piscine

func LastRune(s string) rune {
	count := 0
	for range s {
		count++
	}
	return rune(s[count-1])
}