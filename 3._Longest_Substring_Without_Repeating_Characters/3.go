package longestsubstringwithoutrepeatingcharacters

// 3. Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest
// substring
// without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// a b c d a a

// a b c d a

// left 0

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			l = max(l, index)
		}
		res = max(res, r-l+1)
		m[s[r]] = r + 1
	}

	return res
}
