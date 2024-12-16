package stringtointegeratoi

func myAtoi(s string) int {

	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			length++
			continue
		}
		if s[i] != '-' {
			length++
			continue
		}
	}

	return 0
}
