//This problem was asked by Twitter.
//
//Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.
//
//For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
//
//Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

package problem

type TrieNode struct {
	term     string
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrie() *TrieNode {
	return &TrieNode{
		term:     "",
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

func newNode(term string) *TrieNode {
	return &TrieNode{
		term:     term,
		children: make(map[rune]*TrieNode),
	}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = newNode(node.term + string(c))
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (t *TrieNode) Search(word string) []string {
	node := t

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return nil
		}
		node = node.children[c]
	}

	return grabWords(node)
}

func grabWords(t *TrieNode) []string {
	if len(t.children) == 0 {
		return []string{t.term}
	}

	var results []string

	if t.isWord {
		results = append(results, t.term)
	}

	for _, child := range t.children {
		results = append(results, grabWords(child)...)
	}

	return results
}
