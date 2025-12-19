package main

type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &WordDictionary{}
		}
		curr = curr.children[idx]
	}

	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this, word)
}

func (this *WordDictionary) search(node *WordDictionary, prefix string) bool {
	if node == nil {
		return false
	}

	for i := range prefix {
		if prefix[i] == '.' {
			for j := range node.children {
				if this.search(node.children[j], prefix[i+1:]) {
					return true
				}
			}

			return false
		}

		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return false
		}

		node = node.children[idx]
	}

	return node.end
}
