package bst

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Tree[T constraints.Ordered] struct {
	root *Node[T]
}

type Node[T constraints.Ordered] struct {
	Datum T
	Left  *Node[T]
	Right *Node[T]
}

func New[T constraints.Ordered]() Tree[T] {
	return Tree[T]{root: nil}
}

func (b Tree[T]) String() string {
	var str string
	b.traverse(&str, b.root)
	str = strings.TrimSpace(str)
	if str == "" {
		return "nil"
	}
	return str
}

func (b *Tree[T]) Insert(data ...T) {
	for _, datum := range data {
		b.insertOne(datum, b.root)
	}
}

func (b Tree[T]) Search(datum T) bool {
	if b.root == nil {
		return false
	}
	return b.root.search(datum)
}

func (n Node[T]) search(datum T) bool {
	if datum < n.Datum {
		if n.Left == nil {
			return false
		}
		return n.Left.search(datum)
	}
	if datum > n.Datum {
		if n.Right == nil {
			return false
		}
		return n.Right.search(datum)
	}
	return true
}

func (b Tree[T]) traverse(str *string, n *Node[T]) {
	if n == nil {
		return
	}
	*str += fmt.Sprintf("%v -> %s\n", n.Datum, n.asSet())
	b.traverse(str, n.Left)
	b.traverse(str, n.Right)
}

func (n Node[T]) asSet() string {
	left := "nil"
	if n.Left != nil {
		left = fmt.Sprint(n.Left.Datum)
	}
	right := "nil"
	if n.Right != nil {
		right = fmt.Sprint(n.Right.Datum)
	}
	return fmt.Sprintf("(%s, %s)", left, right)
}

func (b *Tree[T]) insertOne(datum T, n *Node[T]) {
	if b.root == nil {
		b.root = &Node[T]{Datum: datum}
		return
	}

	if datum < n.Datum {
		if n.Left == nil {
			n.Left = &Node[T]{Datum: datum}
			return
		}
		b.insertOne(datum, n.Left)
		return
	}

	if n.Right == nil {
		n.Right = &Node[T]{Datum: datum}
		return
	}
	b.insertOne(datum, n.Right)
}
