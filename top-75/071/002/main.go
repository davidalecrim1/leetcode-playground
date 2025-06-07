package main

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
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

	curr.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root

	for _, c := range word {
		if curr.children == nil {
			return false
		}

		if _, ok := curr.children[c]; !ok {
			return false
		}

		curr = curr.children[c]
	}

	return curr.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root

	for _, c := range prefix {
		if curr.children == nil {
			return false
		}

		if _, ok := curr.children[c]; !ok {
			return false
		}

		curr = curr.children[c]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
