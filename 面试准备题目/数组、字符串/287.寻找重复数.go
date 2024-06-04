package main

/*
初始化指针：

使用两个指针，tortoise 和 hare。初始时都指向数组的第一个元素。
tortoise 一次走一步，hare 一次走两步。
寻找相遇点：

通过移动指针，寻找两个指针相遇的点。这一步保证存在相遇点，因为数组中有重复元素，因此必定存在环。
找到环的入口：

将其中一个指针重置到起始位置，然后两个指针一次都走一步，直到它们再次相遇。相遇点即为环的入口，也就是重复的数。
 */

func findDuplicate(nums []int) int {
	// Step 1: Initialize the tortoise and hare
	tortoise := nums[0]
	hare := nums[0]

	// Step 2: Find the intersection point of the two runners
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	// Step 3: Find the entrance to the cycle
	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return hare
}
