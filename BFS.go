package main

func BFS(puzzleTable EightTable, matchTable EightTable) []Moves {
	queue := make([]EightTable, 0)
	puzzleTable.MoveHistory = nil
	queue = append(queue, puzzleTable)
	var movePath []Moves
	for len(queue) != 0 {
		// dequeue
		u := queue[0]
		queue = queue[1:]

		// for _, e := range u.MoveHistory {
		// 	fmt.Println(len(u.MoveHistory), e.Print())
		// }

		if u.Table == matchTable.Table {
			movePath = u.MoveHistory
			break
		}

		moves := u.Movables()
		if moves.Down {
			n := u.MoveDown()
			queue = append(queue, n)
		}
		if moves.Left {
			n := u.MoveLeft()
			queue = append(queue, n)
		}
		if moves.Right {
			n := u.MoveRight()
			queue = append(queue, n)
		}
		if moves.Up {
			n := u.MoveUp()
			queue = append(queue, n)
		}
	}
	return movePath
}
