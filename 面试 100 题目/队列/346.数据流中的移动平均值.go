// -*- coding:utf-8 -*-
// @Time : 2024/5/7 01:06
// @Author: MJ-CJM
// @File : leetcode/346.数据流中的移动平均值
package main

type MovingAverage struct {
	size    int
	sum     int
	q       []int
}


func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}


func (this *MovingAverage) Next(val int) float64 {
	if len(this.q) == this.size {
		this.sum -= this.q[0]
		this.q = this.q[1:]
	}
	this.sum += val
	this.q = append(this.q, val)
	return float64(this.sum) / float64(len(this.q))
}
