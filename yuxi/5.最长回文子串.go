package yuxi

func longestPalindrome(s string) string {
	length := len(s)
	var dp [1000][1000]bool
	if length <= 1 {
		return s
	}

	for i := 0; i < length; i++ {
		dp[i][i] = true //单个字母是回文
	}
	max := string(s[0])

	//记录从i到j个位置是否是回文串
	for i := length; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			//相邻
			if j == i+1 {
				//相等
				if s[j] == s[i] {
					dp[i][j] = true
				}
			} else {
				if dp[i+1][j-1] == true {
					if s[i] == s[j] {
						dp[i][j] = true
					}
				}
			}

			if dp[i][j] == true {
				if len(max) < j+1-i {
					max = string(s[i : j+1])
				}
			}
		}
	}
	return max

}
