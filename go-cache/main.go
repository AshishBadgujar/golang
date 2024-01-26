package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	//empty queue initialization
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

// least recently used values cache: recently used values should be in cache
func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length -= 1

	delete(c.Hash, n.Val) //deleting from the hash
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n

	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left) //remove last element
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)

		if i < q.Length-1 {
			fmt.Printf("<-->")
		}

		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()

	for _, word := range []string{"parrot", "ashish", "bhargav", "potato", "tree", "dog", "ashish"} {
		cache.Check(word)
		cache.Display()
	}
}

// START CACHE
// add: parrot
// 1 - [{parrot}]
// add: ashish
// 2 - [{ashish}<-->{parrot}]
// add: bhargav
// 3 - [{bhargav}<-->{ashish}<-->{parrot}]
// add: potato
// 4 - [{potato}<-->{bhargav}<-->{ashish}<-->{parrot}]
// add: tree
// 5 - [{tree}<-->{potato}<-->{bhargav}<-->{ashish}<-->{parrot}]
// add: dog
// remove: parrot
// 5 - [{dog}<-->{tree}<-->{potato}<-->{bhargav}<-->{ashish}]
// remove: ashish
// add: ashish
// 5 - [{ashish}<-->{dog}<-->{tree}<-->{potato}<-->{bhargav}]
