package linkedlist

import (
	"testing"
)

func Test (t *testing.T) {
	n1 := Node{"a", nil}
	n2 := Node{"b", nil}
	n3 := Node{"c", nil}
	n4 := Node{"d", nil}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	linked := LinkedList{n1}

	next := Node{"next", nil}

	linked.Push(&next)
	last := &linked.Head
	for last.Next != nil {
		last = last.Next
	}
	if last.val != "next" {
		t.Errorf("Expected 'next' to be the last item, got %s", last.val)
	}

	popped := linked.Pop()
	last = &linked.Head
	for last.Next != nil {
		last = last.Next
	}
	if popped.val != "next" {
		t.Errorf("Expected 'next' to be popped, got %s", popped.val)
	}
	if last.val != "d" {
		t.Errorf("Expected 'd' to be the new last item, got %s", last.val)
	}

	linked.Unshift(&popped)
	if (linked.Head.val != "next") {
		t.Errorf("Expected 'next' to be the new head item, got %s", linked.Head.val)
	}
}