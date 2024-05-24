// -*- coding:utf-8 -*-
// @Time : 2024/5/4 00:22
// @Author: MJ-CJM
// @File : leetcode/340.至多包含 k 个字符的最长子串
package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	//这个就是很典型的滑动窗口题目，答题模板也可以从这个题进行总结
	hasMap :=map[byte]int{}
	maxs := 0
	//分别代表左，右窗口的左右边界
	left,right :=0,0
	for right <len(s){
		//出现次数肯定想到哈希表
		hasMap[s[right]]++
		right++
		//哈希表存入的元素大于k的时候就开始减，减到0就删除，然后缩短左边界
		for len(hasMap) >k{
			//删除左边界元素
			hasMap[s[left]]--
			if 0 ==hasMap[s[left]] {
				delete(hasMap,s[left])
			}
			left++
		}
		//right-left就是最长的大小字串，
		//这个maxs是动态变化的，始终保持的是最大字串的长度
		maxs =max(maxs,right-left)
	}
	return maxs
}

func max(a,b int)int{
	if a>b {
		return a
	}else{
		return b
	}
}

