package main

// Description: This file contains the TrieNode struct which is used to create a Trie data structure.
/**
* Basic concepts
* Trie is a tree-like data structure that stores a dynamic set of strings.It is used to store a dynamic set of strings where the keys are usually strings.
* It is used to search for a string in O(L) time complexity where L is the length of the string.
*
* 1. Node: Each node in the trie contains a map of children nodes and a boolean value to indicate if the node is the end of a word.
 */

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

/**
 * Create a new TrieNode
 */
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
}

// Insert a word into the trie
// Insert adds a word to the Trie. It iterates through each character of the word,
// creating new nodes as necessary if the character is not already present in the
// current node's children. Once all characters are added, it marks the end of the
// word in the Trie.
//
// Parameters:
//   word - the string to be added to the Trie
func (t *Trie) Insert(word string) {
	// Start at the root node
	node := t.Root
	for _, char := range word {

		// If the character is not in the children map, create a new node
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// Search finds all words in the Trie that start with the given prefix
func (t *Trie) Search(prefix string) []string {
	var results []string
	node := t.Root
	for _, char := range prefix {
		if _, exists := node.Children[char]; !exists {
			return results
		}
		node = node.Children[char]
	}
	t.collect(node, prefix, &results)
	return results
}

// collect recursively collects all words in the Trie that start with the given prefix
func (t *Trie) collect(node *TrieNode, prefix string, results *[]string) {
	if node.IsEnd {
		*results = append(*results, prefix)
	}
	for char, child := range node.Children {
		t.collect(child, prefix+string(char), results)
	}
}
