package main

import "fmt"

/*
寻找第一个下降点：

从右向左遍历数组，找到第一个满足 nums[i] < nums[j] 的位置 i，其中 j = i + 1。
这一步确保了 i 之后的子数组是一个非递增序列（即从右到左递增）。

寻找需要交换的另一个位置：

如果找到了第一个下降点（即 i >= 0），则在 i 之后找到一个比 nums[i] 大的最小的元素 nums[k]。
k 初始化为数组末尾，从右向左遍历，找到第一个 nums[k] > nums[i] 的位置。
交换 nums[i] 和 nums[k]，这一步保证新的排列比当前的排列大。

反转 i 之后的元素：

反转从 j（即 i + 1）到数组末尾的子数组，使其成为最小的排列。
这一步确保生成的排列是大于当前排列的最小排列。

 */
func nextPermutation(nums []int)  {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j, k := n-2, n-1, n-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		swap_2(nums, i, k)
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		swap_2(nums, i, j)
	}
	return
}

func swap_2(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{1, 2, 3}
	swap(nums, 0, 1)
	fmt.Println(nums)
}