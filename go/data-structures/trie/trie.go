package trie

import (
	"fmt"
	"strings"
)

type Node struct {
	next  [26]*Node
	isEnd bool
}

type T struct {
	root *Node
}

func New() T {
	return T{
		root: new(Node),
	}
}

func (t T) String() string {
	var str string
	t.root.traverse(&str)
	return str
}

func (t *T) Insert(word string) error {
	word = strings.ToLower(word)
	if !isValid(word) {
		return fmt.Errorf("word should only contain characters a-z")
	}
	n := t.root
	for idx, c := range word {
		i := c - 97
		if n.next[i] == nil {
			n.next[i] = new(Node)
		}
		n.next[i].isEnd = idx == len(word)-1
		n = n.next[i]
	}
	return nil
}

func (t T) Search(word string) bool {
	word = strings.ToLower(word)
	if !isValid(word) {
		return false
	}
	return t.root.search("", word)
}

func (n Node) search(prefix, word string) bool {
	for idx, i := range n.next {
		if i == nil {
			continue
		}
		newPrefix := prefix + string(rune(idx+97))
		if !strings.HasPrefix(word, newPrefix) {
			continue
		}
		if i.isEnd && newPrefix == word {
			return true
		}
		return i.search(newPrefix, word)
	}
	return false
}

func (n Node) traverse(prefix *string) {
	var str string
	for idx, i := range n.next {
		if i == nil {
			continue
		}
		newPrefix := *prefix + string(rune(idx+97))
		i.traverse(&newPrefix)
		str += newPrefix
		if i.isEnd {
			str += "\n"
		}
	}
	if str != "" {
		*prefix = str
	}
}

func isValid(word string) bool {
	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			continue
		}
		return false
	}
	return true
}
