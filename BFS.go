package main

import (
	"fmt"
	"log"
)

func BFS(puzzleTable EightTable, matchTable EightTable) []Moves {
	queue := make([]EightTable, 0)
	puzzleTable.MoveHistory = nil
	firstTable := puzzleTable
	queue = append(queue, puzzleTable)
	var movePath []Moves
	for len(queue) != 0 {
		// log.Println("queue size: ", len(queue))
		// dequeue
		u := queue[0]
		queue = queue[1:]
		// log.Println("queue size: ", len(queue))
		log.Println("current history:")
		for i, e := range u.MoveHistory {
			fmt.Println(i, e.Print())
		}

		if u.Table == matchTable.Table {
			movePath = u.MoveHistory
			break
		}

		moves := u.Movables()
		if moves.Down {
			n := u.Self()
			err := n.MoveDown()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Left {
			n := u.Self()
			err := n.MoveLeft()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Right {
			n := u.Self()
			err := n.MoveRight()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Up {
			n := u.Self()
			err := n.MoveUp()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
	}
	return movePath
}
