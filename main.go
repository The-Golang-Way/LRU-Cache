package main

import (
	"fmt"
)

type Node struct{
	Val string
	Left *Node
	Right *Node
}

type Queue struct{
	Head *Node
	Tail *Node
	Length int
}

type Cache struct{
	Queue Queue
	Hash Hash
}

type Hash map[string] * Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

  return Queue{Head:head, Tail:tail}
}

func (c *Cache) Check(str string){
  node := &Node{}
  
    if val, ok := c.Hash[str]; ok{
      node = c.Remove(val)
    } else {
      node = &Node{Val: str}
    }
    c.Add(node)
  }
}

func (c *Cache) Remove(n *Node) *Node{
  fmt.Printf("remove: %s\n", n.Val)
  left := n.Left
  right := n.Right

  left.Right = right
  right.Left = left
  delete(c.Hash, n.Val)
  return n
}

func (c *Cache) Add(n *Node){
  fmt.Printf("add: %s\n", n.Val)
  tmp := c.Queue.Head.Right
  
  c.Queue.Head.Right = n
  n.Left = c.Queue.Head
  n.Right = tmp
  tmp.Left = n

}

func main() {
	fmt.Println("start the cache baby")
  cache := NewCache()
  for _, turtles := range []string{"leo","ralph","mickey","don"}{
    cache.Check(turtles)
  }
}
