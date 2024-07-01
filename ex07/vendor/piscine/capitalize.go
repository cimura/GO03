package	piscine

func Isnum(r rune) bool {
	if r <= '9' && r >= '0' {
		return true
	}
	return false
}

func Isupper(r rune) bool {
	if (r <= 'Z' && r >= 'A') {
		return true
	}
	return false
}

func Islower(r rune) bool {
	if (r <= 'z' && r >= 'a') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	srune := [] rune(s)
	i := 0
	flag := true
	for range s {
		if (flag == true) {
			if (Islower(srune[i])) {
				srune[i] -= 32
			}
			flag = false
		} else {
			if (Isupper(srune[i])) {
				srune[i] += 32
				flag = false
			}
		}
		if (!Isupper(srune[i]) && !Islower(srune[i]) && !Isnum(srune[i])) {
			flag = true
		}
		i++
	}
	return string(srune)
}