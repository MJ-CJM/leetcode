package main

//type bytes []byte
////自定义排序规则
//func (b bytes) Len() int{
//	return len(b)
//}
//
//func (b bytes)Less(i, j int)bool {
//	return b[i] < b[j]
//}
//
//func (b bytes)Swap(i, j int){
//	b[i], b[j] = b[j], b[i]
//}
//
//func groupAnagrams(strs []string) [][]string {
//	ret := [][]string{}
//	m := make(map[string]int)
//	for _, str := range strs{
//		kbyte := bytes(str)
//		sort.Sort(kbyte)
//		k := string(kbyte)
//		if idx, ok := m[k]; !ok{
//			m[k] = len(ret)
//			ret := append(ret, []string{str})
//		}else{
//			ret[idx] = append(ret[idx],str)
//		}
//	}
//	return ret
//}

// 暴力解法
//func groupAnagrams(strs []string) [][]string {
//	result := [][]string{}
//	n := len(strs)
//	for i := 0; i < n-1; i++ {
//		if strs[i] == "0"{
//			continue
//		}
//		temp := []string{strs[i]}
//		flag := 0
//		for j := i+1; j < n; j++ {
//			if strs[j] == "0"{
//				continue
//			}
//			if j == n-1{
//				if isAnagram(strs[i], strs[j]){
//					flag = 1
//					temp = append(temp, strs[j])
//					strs = append(append(strs[:j], "0"))
//				}
//			}else{
//				if isAnagram(strs[i], strs[j]){
//					flag = 1
//					temp = append(temp, strs[j])
//					strs = append(append(strs[:j], "0"), strs[j+1:]...)
//				}
//			}
//		}
//		if flag == 1{
//			strs = append(append(strs[:i], "0"), strs[i+1:]...)
//		}
//		result = append(result, temp)
//	}
//	for _, v := range strs{
//		temp := []string{}
//		if v != "0"{
//			temp = append(temp, v)
//		}
//		result = append(result, temp)
//	}
//	return result
//}
//
//func isAnagram(s string, t string) bool {
//	s1 := [26]int{}
//	s2 := [26]int{}
//	for _, v := range s {
//		s1['z' - v] += 1
//	}
//	for _, v := range t {
//		s2['z' - v] += 1
//	}
//	return s1 == s2
//}

// 排序综合解法
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	m := make(map[[26]int][]string)
	for _, v := range strs{
		k := getkey(v)
		_, ok := m[k]
		if ok {
			m[k] = append(m[k], v)
		}else{
			m[k] = []string{v}
		}
	}
	for _, v := range m{
		result = append(result, v)
	}
	return result
}

// string-key
func getkey(s string) [26]int{
	res := [26]int{}
	for _, v := range s{
		res[v-'a']++
	}
	return res
}