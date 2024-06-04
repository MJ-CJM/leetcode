package main

/*
初始化：low 和 current 指向 0，high 指向 n-1。
遍历数组：
当 current 指针小于等于 high 指针时，进行以下操作：
如果 nums[current] 是 0（红色），交换 nums[current] 和 nums[low]，然后 current 和 low 指针都向右移动一位。
如果 nums[current] 是 1（白色），不做交换，直接将 current 指针向右移动一位。
如果 nums[current] 是 2（蓝色），交换 nums[current] 和 nums[high]，然后将 high 指针向左移动一位。
 */

func sortColors(nums []int) {
	low, current, high := 0, 0, len(nums)-1

	for current <= high {
		if nums[current] == 0 {
			nums[current], nums[low] = nums[low], nums[current]
			low++
			current++
		} else if nums[current] == 1 {
			current++
		} else {
			nums[current], nums[high] = nums[high], nums[current]
			high--
		}
	}
}
