package plaindromenumber

// 1221
// true

// 1 .
// 2 .
// 22

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	div := 1
	for x >= div*10 {
		div = div * 10
	}

	for x > 0 {
		// ex 1221
		// 1221 / 1000 is 1 => left side
		// 1221 % 10 is 1 => right side
		if x%10 != x/div {
			return false
		}

		// eliminate the x with from outer number

		// x = 1221 % 1000
		// x = 221
		x = x % div
		// x = 221 / 10
		// x = 22
		x = x / 10

		// decrease the div with 100 or two digit
		div /= 100
	}

	return true
}
