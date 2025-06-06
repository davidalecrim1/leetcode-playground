package main

// Time Complexity: O(26) = O(1)
// Space Complexity: O(1)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
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
	node := this.traverse(word)
	return node != nil && node.isEndOfWord
}

func (this *Trie) traverse(s string) *TrieNode {
	curr := this.root

	for _, c := range s {
		if curr.children == nil {
			return nil
		}
		if next, ok := curr.children[c]; ok {
			curr = next
		} else {
			return nil
		}
	}

	return curr
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.traverse(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
