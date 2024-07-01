package piscine

func StrLen(s []string) int {
	count := 0
	i := 0
	for range s {
		for range s[i] {
			count++
		}
		i++
	}
	return count
}

func Join(strs []string, sep string) string {
	var join string
	i := 0
	for range strs {
		join += strs[i]
		join += sep
		i++
	}
	return join[:StrLen(strs)+i-1]
}