package main

import "fmt"

// 纯BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 迭代wordList，把每个word都添加到匹配索引的列表中
	wdict := make(map[string][]int)
	for id, w := range wordList {
		for i := range w {
			k := w[0:i] + "*" + w[i+1:]
			if _, ok := wdict[k]; !ok {
				wdict[k] = []int{}
			}
			wdict[k] = append(wdict[k], id)
		}
	}
	// 建立一个数组，标识一下用过的word；顺手初始化BFS的当前层数为1
	used, l := make([]bool, len(wordList)), 1
	// 偷偷把beginWord放到wordList最后，这样就可以方便的使用下标索引到它
	wordList = append(wordList, beginWord)
	// 来吧，把beginWord放到BFS第一层，开始狂风一般的BFS吧！
	q := []int{len(wordList) - 1}
	for len(q) > 0 { // BFS套路第一招，无脑循环到queue为空方才罢休
		nextQ := []int{}        // BFS套路第二招，刚进入循环就为下一层的迭代做好准备
		l++                     // 这里投机取巧把层数+1
		for _, qid := range q { // BFS套路第三招： 迭代当前层的所有queue里面节点
			w := wordList[qid] // 把这个word给我取出来
			for i := range w { // 找出这个word能变成的所有字典索引
				k := w[0:i] + "*" + w[i+1:]    //生成这个索引
				for _, wid := range wdict[k] { // 在这个索引里面找到跟这个word差了一个字母的那个word
					if wordList[wid] == endWord { // BFS套路第四招： 满足条件就退出吧
						return l
					}
					if !used[wid] { // 要是这个词没用过
						used[wid] = true           // 那好，你现在被用过了
						nextQ = append(nextQ, wid) // 把这个词放到下一层要迭代的queue里面去！
					}
				}
			}
		}
		q = nextQ // BFS套路第五招，既然本层没找到，那么准备迭代下一层吧
	}
	return 0 // 一无所获
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

