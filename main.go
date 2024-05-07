package main

import(
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

func NewCache() Cache{
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue{
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head
}

type Hash map[string]*Node

func main(){
	fmt.Println("START CACHE")
	cache := NewCache()
	for _,word := range []string{"parrot","avocardo","dragonfruit","tree","potato","tomato","tree","dog"}{
		cache.Check(word)
		cache.Display()
	}
}