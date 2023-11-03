package day_heap

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children:    make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}

	node.isEndOfWord = true
}

func (t *Trie) Delete(word string) {
	t.deleteRecursively(t.Root, word, 0)
}

func (t *Trie) deleteRecursively(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEndOfWord {
			return false
		}

		node.isEndOfWord = false
		return len(node.children) == 0
	}

	char := rune(word[index])
	if child, ok := node.children[char]; !ok {
		//word not found
		return false
	} else {
		shouldDeleteChild := t.deleteRecursively(child, word, index+1)
		if shouldDeleteChild {
			delete(node.children, char)
			return len(node.children) == 0
		}
		return false
	}
}
