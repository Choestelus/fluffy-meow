package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func init() {
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(44)
}

func main() {
	fmt.Printf("Creating New Initial Table\n")
	et := NewTable()
	et.Print()
	fmt.Printf("Shuffling...\n")
	if len(os.Args) == 2 {
		count, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("try insert a number and try again")
		}
		et.Shuffle(count)
	} else {
		et.Shuffle(10)
	}
	et.Print()
	solution := BFS(et, NewTable())
	for i, e := range solution {
		if i == 0 {
			// continue
		}
		fmt.Println(i, "move: ", e.Print())
	}
}
