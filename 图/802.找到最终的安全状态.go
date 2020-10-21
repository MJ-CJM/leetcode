package main

import "sort"

//我们将所有的边反向，得到反向图 rgraph，随后将 rgraph 中所有没有入边的节点加入队列中。每一次我们取出队列中的一个节点 u，将它从图中删除，如果此时
//某个节点 v 存在从 u 到 v 的一条边，并且在删掉了这条边后，v 变成了没有入边的节点，那么就把 v 加入队列。以此类推，直到队列为空。最后所有加入过队列
//的节点即为安全的节点。

// 拓扑排序
func eventualSafeNodes(graph [][]int) []int {
	rgraph := make([][]int, len(graph))
	m := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		m[i] = len(graph[i])
		for j := 0; j < len(graph[i]); j++ {
			rgraph[graph[i][j]] = append(rgraph[graph[i][j]], i)
		}
	}
	nums := make([]int, 0)
	for i := 0; i < len(rgraph); i++ {
		if m[i] == 0{
			nums = append(nums, i)
		}
	}
	ans := make([]int, 0)
	for len(nums) > 0 {
		k := nums[0]
		ans = append(ans, k)
		for _, v := range rgraph[k] {
			m[v]--
			if m[v] == 0 {
				nums = append(nums, v)
			}
		}
		nums = nums[1:]
	}
	sort.Ints(ans)
	return ans
}
