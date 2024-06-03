package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1 := make(map[byte]int)
	hash2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash1[s[i]] = i
		hash2[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		if hash1[s[i]] != hash2[t[i]] {
			return false
		}
	}
	return true
}