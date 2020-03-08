package linkedlist

import (
	"sort"
	"testing"
)

func TestLinkedList(t *testing.T) {
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

func TestInsertAfter(t *testing.T) {
	n1 := Node{1, nil}
	n2 := Node{2, nil}
	n3 := Node{3, nil}
	n4 := Node{4, nil}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	linked := LinkedList{n1}

	linked.InsertAfter(&n2, &Node{100, nil})

	last := &linked.Head
	for last != nil {
		if last.val == 2 && last.Next.val != 100 {
			t.Errorf("Expected 100 to come after 2, got %d\n", last.Next.val)
		}
		last = last.Next
	}
}

func TestMerge(t *testing.T) {
	n1, n2, n3, n4 := Node{1, nil}, Node{3, nil}, Node{4, nil}, Node{7, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	l1 := LinkedList{n1}

	m1, m2, m3, m4 := Node{1, nil}, Node{2, nil}, Node{5, nil}, Node{6, nil}
	m1.Next = &m2
	m2.Next = &m3
	m3.Next = &m4
	l2 := LinkedList{m1}

	l3 := Merge(&l1, &l2)
	last := &l3.Head
	merged := make([]int, 8)
	i := 0
	for last != nil {
		merged[i] = last.val
		last = last.Next
		i++
	}
	if !sort.IntsAreSorted(merged) {
		t.Errorf("Expected result to be sorted, got %v", merged)
	}
}
