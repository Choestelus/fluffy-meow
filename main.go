package main

import (
	//"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	et := NewTable()
	et.Print()
	et.Shuffle(1000)
	et.Print()
}
