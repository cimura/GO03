package piscine

func strlen (s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Compare(a, b string) int {
	i := 0
	for range a {
		if (i > strlen(a)-1) {
			return -1
		} else if (i > strlen(b)-1) {
			return 1
		}
		if (rune(a[i]) > rune(b[i])) {
			return 1
		} else if (rune(a[i]) < rune(b[i])) {
			return -1
		}
		i++
	}
	return 0
}