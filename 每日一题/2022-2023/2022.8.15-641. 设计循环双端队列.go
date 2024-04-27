// -*- coding:utf-8 -*-
// @Time : 2022/8/15 11:13 PM
// @Author: MJ-CJM
// @File : leetcode/2022.8.15-641. 设计循环双端队列
package main

type MyCircularDeque struct {
	front, rear int
	elements    []int
}

func constructor(k int) MyCircularDeque {
	return MyCircularDeque{elements: make([]int, k+1)}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.front = (q.front - 1 + len(q.elements)) % len(q.elements)
	q.elements[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear - 1 + len(q.elements)) % len(q.elements)
	return true
}

func (q MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear-1+len(q.elements))%len(q.elements)]
}

func (q MyCircularDeque) IsEmpty() bool {
	return q.rear == q.front
}

func (q MyCircularDeque) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}
