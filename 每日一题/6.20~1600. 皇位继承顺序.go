// -*- coding:utf-8 -*-
// @Time : 2021/6/20 6:58 下午
// @Author: MJ-CJM
// @File : leetcode/6.20~1600. 皇位继承顺序
package main

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) (t ThroneInheritance) {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (t *ThroneInheritance) Birth(parentName, childName string) {
	t.edges[parentName] = append(t.edges[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var preorder func(string)
	preorder = func(name string) {
		if !t.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range t.edges[name] {
			preorder(childName)
		}
	}
	preorder(t.king)
	return
}


/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
