package linkedlist

type Node struct {
	val  int
	Next *Node
}

type LinkedList struct {
	Head Node
}

func (l *LinkedList) Push(node *Node) {
	last := &l.Head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = node
}

func (l *LinkedList) Pop() Node {
	last := &l.Head
	for last.Next.Next != nil {
		last = last.Next
	}
	popped := last.Next
	last.Next = nil
	return *popped
}

func (l *LinkedList) Unshift(newHead *Node) {
	prev := l.Head
	newHead.Next = &prev
	l.Head = *newHead
}

func (l *LinkedList) Shift() Node {
	last := l.Head
	l.Head = *l.Head.Next
	return last
}

func (l *LinkedList) InsertAfter(ref *Node, newNode *Node) {
	if ref.Next == nil {
		ref.Next = newNode
	} else {
		tmp := ref.Next
		ref.Next = newNode
		ref.Next.Next = tmp
	}
}

func Merge(l1 *LinkedList, l2 *LinkedList) LinkedList {
	start := Node{-1, nil}
	last := &start
	i, j := &l1.Head, &l2.Head
	for i != nil && j != nil {
		if i.val < j.val {
			last.Next = i
			last = i
			i = i.Next
		} else {
			last.Next = j
			last = j
			j = j.Next
		}
	}

	if i != nil {
		last.Next = i
	}

	if j != nil {
		last.Next = j
	}

	return LinkedList{*start.Next}
}
