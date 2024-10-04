package trie

import (
	"testing"
)

func TestAddValue(test *testing.T) {
	trie := NewTrie()
	trie.
		AddValue("apple").
		AddValue("banana").
		AddValue("bandana")

	// There should be 2 top-level nodes, one for "a" and one for "b"
	if trie.Size() != 2 {
		test.Error()
	}
}

func TestAddValuesFromArray(test *testing.T) {
	trie := NewTrie()
	trie.AddValuesFromArray([]string{"apple", "banana", "bandana", "cherry"})

	// There should be 3 top-level nodes, one for "a", one for "b" and one for "c"
	if trie.Size() != 3 {
		test.Error()
	}
}

func TestSearch(test *testing.T) {
	trie := NewTrie()
	trie.AddValuesFromArray([]string{"apple", "banana", "bandana", "cherry"})

	node1 := trie.Search("apple")
	node2 := trie.Search("banana")
	node3 := trie.Search("bandana")
	node4 := trie.Search("cherry")
	node5 := trie.Search("coconut")

	// node5 should be nil because the word "coconut" will not be mapped to any nodes
	if node1 == nil || node2 == nil || node3 == nil || node4 == nil || node5 != nil {
		test.Error()
	}
}

func TestSearchMatch(test *testing.T) {
	trie := NewTrie()
	trie.AddValuesFromArray([]string{"apple", "banana", "bandana", "cherry"})

	node1 := trie.Search("app")
	node2 := trie.Search("appl")
	node3 := trie.Search("apple")

	// node3 should match a full value, the others should not
	if node1.Matched() || node2.Matched() || !node3.Matched() {
		test.Error()
	}
}
