package	piscine

func strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func NRune(s string, n int) rune {
	if n > strlen(s) || n < 0 {
		return 0
	}
	return rune(s[n-1])
}