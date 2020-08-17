package main

func judgePoint24(nums []int) bool {
	// write code here
	target := 24
	iterm := 0
	n := len(nums)
	level := 0
	flag := false
	_dfs_24(nums, target, iterm, level, n, &flag)
	return flag
}

func _dfs_24(arr []int, target int, iterm int, level int, n int, flag *bool) {
	// terminator
	if level >= n {
		if iterm == 24 {
			*flag = true
			return
		}
		return
	}

	// process && drill down
	for i := 0; i < len(arr); i++ {
		c1 := iterm + arr[i]
		c2 := iterm - arr[i]
		c3 := iterm * arr[i]
		tmp := make([]int, len(arr)-1)
		copy(tmp[0:], arr[0:i])
		copy(tmp[i:], arr[i+1:])
		_dfs_24(tmp, target, c1, level+1, n, flag)
		_dfs_24(tmp, target, c2, level+1, n, flag)
		_dfs_24(tmp, target, c3, level+1, n, flag)
		if iterm % arr[i] == 0{
			c4 := iterm / arr[i]
			_dfs_24(tmp, target, c4, level+1, n, flag)
		}
	}
}

