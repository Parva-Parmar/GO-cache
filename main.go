package main

import(
	"fmt"
)

type Node struct{

}


type Queue struct{

}

type Cache struct{
	Queue Queue
	Hash Hash
}

func NewCache() Cache{
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue{
	
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