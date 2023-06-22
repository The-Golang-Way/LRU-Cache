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

	return Queue{}
}

func main() {
	fmt.Println("start the cache baby")
}