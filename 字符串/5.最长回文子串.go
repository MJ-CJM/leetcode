package main

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}
	dp := make([][]bool, length)
	start := 0
	maxlen := 1
	for r := 0; r < length; r++ {
		dp[r] = make([]bool, length)
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			}else{
				dp[l][r] = false
			}
			if dp[l][r] {
				curlen := r-l+1
				if curlen > maxlen {
					maxlen = curlen
					start = l
				}
			}
		}
	}
	return s[start:start+maxlen]
}
