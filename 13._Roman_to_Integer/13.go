package romantointeger

func romanToInt(s string) int {
	res := 0
	length := len(s) - 1
	for i, j := length, length; i >= 0; i-- {
		j = i - 1
		switch string(s[i]) {
		case "I":
			res += 1
		case "V":
			if j > -1 && string(s[j]) == "I" {
				res += 4
				i--
				continue
			}
			res += 5
		case "X":
			if j > -1 && string(s[j]) == "I" {
				res += 9
				i--
				continue
			}
			res += 10
		case "L":
			if j > -1 && string(s[j]) == "X" {
				res += 40
				i--
				continue
			}
			res += 50
		case "C":
			if j > -1 && string(s[j]) == "X" {
				res += 90
				i--
				continue
			}
			res += 100
		case "D":
			if j > -1 && string(s[j]) == "C" {
				res += 40
				i--
				continue
			}
			res += 500
		case "M":
			if j > -1 && string(s[j]) == "C" {
				res += 900
				i--
				continue
			}
			res += 1000
		}
	}

	return res
}
