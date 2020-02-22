package linkedlist

type Node struct {
	val string
	Next *Node
}

type LinkedList struct {
	Head Node
}

func (l *LinkedList) Push(node *Node) {
	last := &l.Head
	for (last.Next != nil) {
		last = last.Next
	}
	last.Next = node
}

func (l *LinkedList) Pop() Node {
	last := &l.Head
	for (last.Next.Next != nil) {
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