package main

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	word     bool
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root

	for _, c := range word {
		if curr.children == nil {
			curr.children = make(map[rune]*TrieNode)
		}
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = &TrieNode{}
		}

		curr = curr.children[c]
	}

	curr.word = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(j int, head *TrieNode) bool {
		curr := head
		for i := j; i < len(word); i++ {
			c := rune(word[i])

			if c == '.' {
				for _, v := range curr.children {
					if dfs(i+1, v) {
						return true
					}
				}

				return false

			} else {
				if _, ok := curr.children[c]; !ok {
					return false
				}
				curr = curr.children[c]
			}
		}
		return curr.word
	}

	return dfs(0, this.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
