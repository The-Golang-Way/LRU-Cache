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
  }
}

func main() {
	fmt.Println("start the cache baby")
  cache := NewCache()
  for _, turtles := range []string{"leo","ralph","mickey","don"}{
    cache.Check(turtles)
  }
}
