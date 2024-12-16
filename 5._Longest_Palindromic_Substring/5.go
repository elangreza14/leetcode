package longestpalindromicsubstring

func longestPalindrome(s string) string {
	res := ""
	lenRes := 0

	for index := range s {
		i, j := index, index
		for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
			if j-i+1 >= lenRes {
				lenRes = j - i + 1
				res = s[i : j+1]
			}

			i--
			j++
		}

		i, j = index, index+1
		for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
			if j-i+1 >= lenRes {
				lenRes = j - i + 1
				res = s[i : j+1]
			}

			i--
			j++
		}
	}

	return res
}
