package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	fmt.Printf("Creating New Initial Table\n")
	et := NewTable()
	et.Print()
	fmt.Printf("Shuffling...\n")
	et.Shuffle(1000)
	et.Print()
}
