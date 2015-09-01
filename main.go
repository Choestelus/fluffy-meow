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
	rand.Seed(45)
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
		et = et.Shuffle(count)
	} else {
		et = et.Shuffle(10)
	}
	et.Print()
	// for _, e := range et.MoveHistory {
	// 	fmt.Println(e.Print())
	// }
	fmt.Println("--------------------------------------------------------------------------------")
	var displayTable EightTable
	for i := 0; i < TABLE_SIZE; i++ {
		for j := 0; j < TABLE_SIZE; j++ {
			displayTable.Table[i][j] = et.Table[i][j]
		}
	}
	solution := BFS(et, NewTable())
	for i, e := range solution {
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println(i+1, "move: ", e.Print())
		if e.Down {
			displayTable = displayTable.MoveDown()
			displayTable.Print()
		}
		if e.Left {
			displayTable = displayTable.MoveLeft()
			displayTable.Print()
		}
		if e.Up {
			displayTable = displayTable.MoveUp()
			displayTable.Print()
		}
		if e.Right {
			displayTable = displayTable.MoveRight()
			displayTable.Print()
		}
	}
}
