package main

func BFS(puzzleTable EightTable, matchTable EightTable) []Moves {
	queue := make([]EightTable, 0)
	queue = append(queue, puzzleTable)
	var movePath []Moves
	for len(queue) != 0 {
		// log.Println("queue size: ", len(queue))
		// dequeue
		u := queue[0]
		queue = queue[1:]
		// log.Println("queue size: ", len(queue))

		if u.Table == matchTable.Table {
			movePath = u.MoveHistory
			break
		}

		moves := u.Movables()
		// log.Println("moveables:", "up", moves.Up, "down", moves.Down, "left", moves.Left, "right", moves.Right)
		if moves.Down {
			// u.Print()
			// log.Println("enqueueing down")
			n := u.Self()
			err := n.MoveDown()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Left {
			// u.Print()
			// log.Println("enqueueing left")
			n := u.Self()
			err := n.MoveLeft()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Right {
			// u.Print()
			// log.Println("enqueueing right")
			n := u.Self()
			err := n.MoveRight()
			if err != nil {
				panic(err)
			}
			queue = append(queue, n)
		}
		if moves.Up {
			// u.Print()
			// log.Println("enqueueing up")
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
