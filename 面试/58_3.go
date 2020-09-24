package main

/**
 * 合并数组
 * @param arrayA int整型一维数组 数组A
 * @param arrayB int整型一维数组 数组B
 * @return int整型一维数组
 */
func mergerArrays( arrayA []int ,  arrayB []int ) []int {
	// write code here
	res := []int{}
	n := len(arrayA)
	m := len(arrayB)
	i, j := 0, 0
	for i < n && j < m {
		if arrayA[i] == arrayB[j] {
			res = append(res, arrayA[i])
			i++
			j++
			continue
		}
		if arrayA[i] > arrayB[j] {
			j++
			continue
		}
		if arrayA[i] < arrayB[j] {
			i++
			continue
		}
	}
	return res
}
