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

type Cache struct{}

type Hash map[string] * Node

func main() {
	fmt.Println("start the cache baby")
}