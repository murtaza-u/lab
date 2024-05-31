package hashmap

import (
	"fmt"
	"strings"
)

const arraySize = 10

type Map struct {
	array  [arraySize]*Node
	hasher Hasher
}

type Node struct {
	k    string
	v    any
	next *Node
}

func New() Map {
	return Map{
		hasher: DefaultHasher(arraySize),
	}
}

func (m Map) String() string {
	var str string
	for _, head := range m.array {
		if head == nil {
			continue
		}
		i := head
		for i != nil {
			str += fmt.Sprintf("%s(%v) -> ", i.k, i.v)
			i = i.next
		}
		str += "nil\n"
	}
	if str == "" {
		str = "nil"
	}
	return strings.TrimSpace(str)
}

func (m *Map) Insert(k string, v any) {
	hash := m.hasher.Hash(k)
	if m.array[hash] == nil {
		m.array[hash] = &Node{k: k, v: v}
		return
	}

	var replace bool
	i := m.array[hash]
	for i.next != nil {
		if i.k == k {
			replace = true
			break
		}
		i = i.next
	}

	if replace {
		i.v = v
		return
	}
	i.next = &Node{k: k, v: v}
}

func (m Map) Get(k string) any {
	hash := m.hasher.Hash(k)
	head := m.array[hash]
	if head == nil {
		return nil
	}
	i := head
	for i != nil {
		if i.k == k {
			return i.v
		}
		i = i.next
	}
	return nil
}

func (m *Map) Delete(k string) {
	hash := m.hasher.Hash(k)
	head := m.array[hash]
	if head == nil {
		return
	}

	if head.k == k {
		m.array[hash] = head.next
		return
	}

	prev := head
	curr := head.next
	for curr != nil {
		if curr.k != k {
			prev = curr
			curr = curr.next
			continue
		}
		prev.next = curr.next
		break
	}
}
