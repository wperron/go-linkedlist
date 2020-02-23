package linkedlist

import (
	"testing"
)

func Test (t *testing.T) {
	n1 := Node{1, nil}
	n2 := Node{2, nil}
	n3 := Node{3, nil}
	n4 := Node{4, nil}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	linked := LinkedList{n1}

	next := Node{100, nil}

	linked.Push(&next)
	last := &linked.Head
	for last.Next != nil {
		last = last.Next
	}
	if last.val != 100 {
		t.Errorf("Expected 'next' to be the last item, got %d", last.val)
	}

	popped := linked.Pop()
	last = &linked.Head
	for last.Next != nil {
		last = last.Next
	}
	if popped.val != 100 {
		t.Errorf("Expected 'next' to be popped, got %d", popped.val)
	}
	if last.val != 4 {
		t.Errorf("Expected 'd' to be the new last item, got %d", last.val)
	}

	linked.Unshift(&popped)
	if linked.Head.val != 100 {
		t.Errorf("Expected 'next' to be the new head item, got %d", linked.Head.val)
	}
}