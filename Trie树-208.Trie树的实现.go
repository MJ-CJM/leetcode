package main

type Trie struct {
	name    string
	subTrie map[string]*Trie
	// tag来标记是否是一个完整的单词
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		name:    "",
		subTrie: make(map[string]*Trie, 0),
		isWord:  false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, v := range word {
		value, ok := root.subTrie[string(v)]
		if ok {
			root = value
		} else {
			newNode := make(map[string]*Trie, 0)
			root.subTrie[string(v)] = &Trie{
				name:    string(v),
				subTrie: newNode,
				isWord:  false,
			}
			root = root.subTrie[string(v)]
		}
	}
	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isWord, ok := this.searchByWord(word)
	if ok && isWord {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.searchByWord(prefix)
	return ok
}

func (this *Trie) searchByWord(word string) (bool, bool) {
	var value *Trie
	var ok bool
	root := this
	for _, v := range word {
		value, ok = root.subTrie[string(v)]
		if ok {
			root = value
			continue
		}
		return false, false
	}
	return value.isWord, true
}
