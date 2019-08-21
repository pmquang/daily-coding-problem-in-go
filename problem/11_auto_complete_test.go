package problem

import (
	"reflect"
	"testing"
)

var trieNodeInsertTests = []struct {
	in  []string
	out []string
}{
	{[]string{"de", "d", ""}, []string{"", "d", "de"}},
	{[]string{"de", "deer", "deal"}, []string{"de", "deer", "deal"}},
	{[]string{"de", "deer", "deal", "de"}, []string{"de", "deer", "deal"}},
}

var trieNodeSearchTests = []struct {
	term string
	dict []string
	out  []string
}{
	{"d", []string{"de", "deer", "deal"}, []string{"de", "deer", "deal"}},
	{"de", []string{"de", "deer", "deal"}, []string{"de", "deer", "deal"}},
	{"dee", []string{"de", "deer", "deal"}, []string{"deer"}},
	{"deel", []string{"de", "deer", "deal"}, nil},
}

func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.term != "" {
		t.Error("Expect root term to be empty.")
	}

	if trie.isWord {
		t.Error("root node can't be a word")
	}

	if len(trie.children) > 0 {
		t.Error("new trie can't have any children")
	}
}

func TestTrieNode_Insert(t *testing.T) {
	for _, tc := range trieNodeInsertTests {
		trie := NewTrie()

		for _, word := range tc.in {
			trie.Insert(word)
		}

		if actual := trie.Search(""); !reflect.DeepEqual(actual, tc.out) {
			t.Errorf("expecting %v, got %v", tc.out, actual)
		}
	}
}

func TestTrieNode_Search(t *testing.T) {
	for _, tc := range trieNodeSearchTests {
		trie := NewTrie()
		for _, word := range tc.dict {
			trie.Insert(word)
		}

		if actual := trie.Search(tc.term); !reflect.DeepEqual(actual, tc.out) {
			t.Errorf("expecting %v, got %v", tc.out, actual)
		}
	}
}
