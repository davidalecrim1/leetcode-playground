package main

type Trie struct {
	children [26]*Trie
	end      bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this

	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}

		curr = curr.children[idx]
	}

	curr.end = true
}

func (this *Trie) Search(word string) bool {
	found, curr := this.exists(word)
	if found {
		return curr.end
	}

	return false
}

func (this *Trie) exists(prefix string) (bool, *Trie) {
	curr := this

	for i := range prefix {
		idx := prefix[i] - 'a'
		if curr.children[idx] == nil {
			return false, nil
		}

		curr = curr.children[idx]
	}

	return true, curr
}

func (this *Trie) StartsWith(prefix string) bool {
	found, _ := this.exists(prefix)
	return found
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
