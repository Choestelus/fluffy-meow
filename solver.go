package main

func BFSSolve(solutionTable EightTable, puzzleTable EightTable, moves []Moves) []Moves {
	if solutionTable == puzzleTable {
		return moves
	}

	m := puzzleTable.Movables()

	if m.Down {
		err := puzzleTable.MoveDown()
		if err != nil {
			panic(err)
		}
		moves = append(moves, Moves{Down: true})
		BFSSolve(solutionTable, puzzleTable, moves)
	}
	if m.Up {
		err := puzzleTable.MoveUp()
		if err != nil {
			panic(err)
		}
		moves = append(moves, Moves{Up: true})
		BFSSolve(solutionTable, puzzleTable, moves)
	}
	if m.Left {
		err := puzzleTable.MoveLeft()
		if err != nil {
			panic(err)
		}
		moves = append(moves, Moves{Left: true})
		BFSSolve(solutionTable, puzzleTable, moves)
	}
	if m.Right {
		err := puzzleTable.MoveRight()
		if err != nil {
			panic(err)
		}
		moves = append(moves, Moves{Right: true})
		BFSSolve(solutionTable, puzzleTable, moves)
	}
}
