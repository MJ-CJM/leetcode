package main

// 转化为岛屿问题
// DFS/BFS
// 并查集
// a. N --> 各自独立集合
// b. 遍历好友关系矩阵M, M[i][j] -- > 合并
// c. 看有多少孤立的集合
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	length := len(M)
	ufind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ { //对称矩阵不用重复计算
			if M[i][j] == 1 {
				ufind.Union(i, j)
			}
		}
	}
	return ufind.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}


// 无递归并查集
func findCircleNum1(M [][]int) int {
	n := len(M)
	p := make([]*[]int, n)
	for i := 0; i < n; i++ {
		p[i] = &[]int{i}
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			i, j := i, j
			if M[i][j] == 1 && p[i] != p[j] {
				if len(*p[i]) < len(*p[j]) {
					i, j = j, i
				}
				*p[i] = append(*p[i], *p[j]...)
				for _, k := range *p[j] {
					p[k] = p[i]
				}
				ans--
			}
		}
	}
	return ans
}

// 深度优先遍历
func findCircleNum2(M [][]int) int {
	ans, v := 0, make([]bool, len(M))
	var f func(int); f = func(i int) {
		for j := range M {
			if M[i][j] == 1 && !v[j] {
				v[j] = true
				f(j)
			}
		}
	}
	for i := range M {
		if !v[i] {
			v[i] = true
			f(i)
			ans++
		}
	}
	return ans
}

// 传统并查集
func findCircleNum3(M [][]int) int {
	n := len(M)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				pi, pj := f(p, i), f(p, j)
				if pi != pj {
					p[pj] = pi
					ans--
				}
			}
		}
	}
	return ans
}

func f(p []int,x int) int {
	if p[x] != x {
		p[x] = f(p, p[x])
	}
	return p[x]
}