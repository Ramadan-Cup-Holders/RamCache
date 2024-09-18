package structure

import (
	"strings"
)

// Node represents a single node in the radix tree
type Node struct {
	Children map[string]*Node // Children nodes
	IsEnd    bool             // Indicates if the node is the end of a word
	Value    string           // Value stored at the node
}

// RadixTree represents the entire radix tree
type RadixTree struct {
	Root *Node
}

// NewRadixTree initializes a new radix tree
func NewRadixTree() *RadixTree {
	return &RadixTree{
		Root: &Node{
			Children: make(map[string]*Node),
		},
	}
}

// Insert adds a word to the radix tree
func (r *RadixTree) Insert(word, value string) {
	node := r.Root
	for len(word) > 0 {
		var prefix string
		for key := range node.Children {
			if strings.HasPrefix(word, key) {
				prefix = key
				break
			}
		}

		// If a prefix exists, traverse deeper
		if prefix != "" {
			node = node.Children[prefix]
			word = strings.TrimPrefix(word, prefix)
		} else {
			// If no prefix exists, insert new node
			node.Children[word] = &Node{
				Children: make(map[string]*Node),
				IsEnd:    true,
				Value:    value,
			}
			return
		}
	}

	// Mark the last node as the end of the word
	node.IsEnd = true
	node.Value = value
}

// Search looks for a word in the radix tree and returns its value
func (r *RadixTree) Search(word string) (string, bool) {
	node := r.Root
	for len(word) > 0 {
		var prefix string
		for key := range node.Children {
			if strings.HasPrefix(word, key) {
				prefix = key
				break
			}
		}

		// If no prefix is found, the word does not exist
		if prefix == "" {
			return "", false
		}

		node = node.Children[prefix]
		word = strings.TrimPrefix(word, prefix)
	}

	// Return the value if it's the end of a word
	if node.IsEnd {
		return node.Value, true
	}

	return "", false
}
