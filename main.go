package main

import (
	"fmt"
	"math/rand"
)

func init() {
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(42)
}

func main() {
	fmt.Printf("Creating New Initial Table\n")
	et := NewTable()
	et.Print()
	fmt.Printf("Shuffling...\n")
	et.Shuffle(100)
	et.Print()
	solution := BFS(et, NewTable())
	for i, e := range solution {
		fmt.Println(i, "move: ", e.Print())
	}
}
