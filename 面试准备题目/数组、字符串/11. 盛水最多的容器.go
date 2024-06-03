package main


func maxArea(height []int) int {
	n := len(height)
	i := 0
	j := n - 1
	res := 0
	for i < j {
		tmp := countArea(i, j, height)
		res = max_count(tmp, res)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func countArea(i, j int, height []int) int {
	x := j - i
	y := 0

	if height[i] < height[j] {
		y = height[i]
	} else {
		y = height[j]
	}

	return x * y
}

func max_count(x, y int) int {
	if x > y {
		return x
	}
	return y
}