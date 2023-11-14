package day_heap

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func (t *Trie) find(prefix string) []string {
	node := t.Root

	// Traverse the Trie to the node corresponding to the prefix
	for _, char := range prefix {
		if nextNode, ok := node.children[char]; ok {
			node = nextNode
		} else {
			// If any character in the prefix is not found, return an empty result
			return []string{}
		}
	}

	// Perform a depth-first search to find all words starting from the given prefix
	result := t.findWordsFromNode(node, prefix)
	return result
}

func (t *Trie) findWordsFromNode(node *TrieNode, currentPrefix string) []string {
	result := []string{}

	// If the current node is the end of a word, add it to the result
	if node.isEndOfWord {
		result = append(result, currentPrefix)
	}

	// Recursively traverse child nodes
	for char, childNode := range node.children {
		nextPrefix := currentPrefix + string(char)
		words := t.findWordsFromNode(childNode, nextPrefix)
		result = append(result, words...)
	}

	return result
}

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	assert.Equal(t, heap.Length, 0)
	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(1000)
	assert.Equal(t, heap.Length, 5)
	assert.Equal(t, heap.Delete(), 1)
	assert.Equal(t, heap.Delete(), 2)
	assert.Equal(t, heap.Length, 3)
}

func TestTries(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")

	// Call the Find method
	result := trie.find("fo")

	// Define the expected result and sort it
	expected := []string{"foo", "fool", "foolish"}
	sort.Strings(expected)

	// Sort the actual result
	sort.Strings(result)

	assert.ElementsMatch(t, expected, result)

	trie.Delete("fool")
	resultAfterDel := trie.find("fo")

	// Call the Find method

	// Define the expected result and sort it
	expectedAfterResult := []string{"foo", "foolish"}
	sort.Strings(expectedAfterResult)

	// Sort the actual result
	sort.Strings(resultAfterDel)

	assert.ElementsMatch(t, expectedAfterResult, resultAfterDel)

}
