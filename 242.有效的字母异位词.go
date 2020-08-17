package main

func isAnagram(s string, t string) bool {
	a := [26]int{}
	b := [26]int{}
	for _, v := range s{
		a['z'-v] += 1
	}
	for _, v := range t{
		b['z'-v] += 1
	}
	return a == b
}


