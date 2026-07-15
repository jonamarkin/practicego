package deque

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Deque struct {
	head  *Node
	tail  *Node
	count int
}

func NewDeque() *Deque {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &Deque{
		head:  head,
		tail:  tail,
		count: 0,
	}
}

func (d *Deque) AddFirst(value int) {
	newNode := &Node{value: value}
	firstRealNode := d.head.next

	newNode.next = firstRealNode
	newNode.prev = d.head

	d.head.next = newNode
	firstRealNode.prev = newNode

	d.count++
}
