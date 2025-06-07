package main

import "fmt"

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type WordDictionary struct {
	root *TrieNode
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

	curr.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this.root

	var dfs func(int, *TrieNode) bool
	dfs = func(j int, head *TrieNode) bool {
		prefix := word[j:]

		for i, char := range prefix {
			if char == '.' {
				for _, next := range head.children {
					if dfs(j+i+1, next) {
						return true
					}
				}

				return false

			} else {
				if _, ok := head.children[char]; !ok {
					return false
				}

				head = head.children[char]
			}
		}

		return head.isEndOfWord
	}

	return dfs(0, curr)
}

func main() {
	dict := Constructor()
	dict.AddWord("bad")

	got := dict.Search("b..")
	fmt.Printf("expected 'true' got: %v\n", got)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
