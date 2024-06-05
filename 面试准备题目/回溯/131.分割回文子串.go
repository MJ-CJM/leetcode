package main

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) (ans [][]string) {
	path := []string{}
	n := len(s)

	// start 表示当前这段回文子串的开始位置
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...)) // 复制 path
			return
		}

		// 不选 i 和 i+1 之间的逗号（i=n-1 时一定要选）
		if i < n-1 {
			dfs(i+1, start)
		}

		// 选 i 和 i+1 之间的逗号（把 s[i] 作为子串的最后一个字符）
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1) // 下一个子串从 i+1 开始
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, 0)
	return
}

