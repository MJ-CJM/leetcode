package main

import (
	"strconv"
	"strings"
)

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// 方法一：广度优先
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	queue := []*TreeNode{root}
	vals := make([]string, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			vals = append(vals, "#")
		}else{
			vals = append(vals, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	if vals[idx] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{
		Val: val,
	}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{
				Val: val,
			}
			node.Left = left
			queue = append(queue, node.Left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{
				Val: val,
			}
			node.Right = right
			queue = append(queue, node.Right)
		}
	}
	return root
}

// 深度优先遍历
func (this *Codec) serialize_2(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize_2(root.Left) + "," + this.serialize_2(root.Right)
}

func dfs(valsPtr *[]string) *TreeNode {
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	if val == "#" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{
		Val: valInt,
	}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)
	return node
}

func (this *Codec) deserialize_2(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}
