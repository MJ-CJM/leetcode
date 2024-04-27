// -*- coding:utf-8 -*-
// @Time : 2022/8/16 11:31 PM
// @Author: MJ-CJM
// @File : leetcode/2022.8.16-1656. 设计有序流
package main

type OrderedStream struct {
	stream []string
	ptr    int
}

func constructor2(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.stream[idKey] = value
	start := s.ptr
	for s.ptr < len(s.stream) && s.stream[s.ptr] != "" {
		s.ptr++
	}
	return s.stream[start:s.ptr]
}
