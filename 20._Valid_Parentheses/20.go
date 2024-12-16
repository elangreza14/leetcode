package validparentheses

func isValid(s string) bool {
	x := []rune{}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			x = append(x, v)
			continue
		case ')':
			if len(x) > 0 && x[len(x)-1] != '(' {
				return false
			}
		case ']':
			if len(x) > 0 && x[len(x)-1] != '[' {
				return false
			}
		case '}':
			if len(x) > 0 && x[len(x)-1] != '{' {
				return false
			}
		}
		if len(x) > 0 {
			x = x[:len(x)-1]
		}
	}

	return len(x) == 0
}
