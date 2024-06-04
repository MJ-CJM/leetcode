package main

/*
使用 Kahn 的算法来进行拓扑排序，通过计算每个节点的入度来判断是否存在环。

初始化：

使用一个数组 inDegree 来记录每个节点的入度。
使用一个邻接表 adjList 来表示图。
构建图：

遍历 prerequisites 数组，更新 inDegree 和 adjList。
BFS：

将所有入度为 0 的节点加入队列。
当队列不为空时，取出一个节点，遍历它的邻居节点，将邻居节点的入度减 1。如果邻居节点的入度变为 0，则将其加入队列。
计数已处理的节点数量。如果已处理的节点数量等于课程数量，则说明可以完成所有课程，否则不能完成。
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化入度数组和邻接表
	inDegree := make([]int, numCourses)
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		inDegree[prereq[0]]++
		adjList[prereq[1]] = append(adjList[prereq[1]], prereq[0])
	}

	// 初始化队列，将所有入度为 0 的节点加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 进行 BFS
	processed := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		processed++
		for _, neighbor := range adjList[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 检查是否所有节点都被处理
	return processed == numCourses
}

/*
使用 DFS 来检测图中是否存在环。如果存在环，则无法完成所有课程。

初始化：

使用一个邻接表 adjList 来表示图。
使用一个数组 visited 来记录每个节点的访问状态：0 表示未访问，1 表示访问中，2 表示已访问。
构建图：

遍历 prerequisites 数组，更新 adjList。
DFS：

对每个节点进行 DFS，如果在 DFS 的过程中遇到一个访问中的节点，说明存在环。
DFS 结束后，如果没有遇到环，说明可以完成所有课程。

 */

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化邻接表
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		adjList[prereq[1]] = append(adjList[prereq[1]], prereq[0])
	}

	// 初始化访问状态数组
	visited := make([]int, numCourses)

	// 定义 DFS 函数
	var dfs func(int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			return false // 检测到环
		}
		if visited[node] == 2 {
			return true // 已经处理过这个节点
		}
		visited[node] = 1 // 标记为访问中
		for _, neighbor := range adjList[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		visited[node] = 2 // 标记为已访问
		return true
	}

	// 对每个节点进行 DFS
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
