package main

import "fmt"

// 迭代
func subsets(nums []int) [][]int {
	var result = [][]int{}
	result = append(result, []int{})
	for _, v := range nums{
		newsets := make([][]int, len(result))
		for k, v1 := range result{
			v1 = append(v1, v)
			newsets[k] = append(newsets[k], v1...)
		}
		result = append(result, newsets...)
	}

	return result
}

// 回溯递归
func subset(nums []int) [][]int{
	result := [][]int{}
	iterm := []int{}
	result = append(result, iterm)
	if nums == nil{
		return result
	}
	dfs(nums, &result, &iterm, 0)
	return result
}

func dfs(nums []int, result *[][]int, iterm *[]int, i int) {
	// terminator
	if i >= len(nums){
		return
	}
	// process(split big problem)
	*iterm = append(*iterm, nums[i])
	temp := make([]int, len(*iterm))
	for i, v := range *iterm{
		temp[i] = v
	} //必须转换一下，不然指针指向的地方会改变
	*result = append(*result, temp)

	// drill down
	dfs(nums, result, iterm, i+1)

	// reverse states
	*iterm = (*iterm)[:len(*iterm)-1]
	dfs(nums, result, iterm, i+1)
	return
}



func main(){
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}