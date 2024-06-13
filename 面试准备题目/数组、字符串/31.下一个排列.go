package main

import "fmt"

/*
从后向前找到第一个下降元素：

从数组末尾开始，向前遍历，找到第一个 nums[i] < nums[i + 1] 的位置 i。这个步骤的目的是找到需要改变的最低位，保证生成的新排列是字典序中的下一个更大值。
从后向前找到第一个大于 nums[i] 的元素：

再次从数组末尾开始，向前遍历，找到第一个大于 nums[i] 的位置 k。这个步骤的目的是找到一个尽量小的数字与 nums[i] 交换，以保证新的排列是紧跟当前排列之后的最小值。
交换 nums[i] 和 nums[k]：

交换 nums[i] 和 nums[k] 的值。这一步确保了前面的部分已经接近目标排列，但仍需要调整以形成下一个字典序。
反转 nums[i+1] 到数组末尾的部分：

最后，反转 nums[i+1] 到数组末尾的部分，使其成为升序排列。这是因为在交换后，nums[i+1] 到数组末尾的部分依旧是降序排列，而我们需要最小的字典序排列，所以要将其反转成升序。

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