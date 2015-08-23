package main

func (et EightTable) BFSSolve() {
	moves := et.Movables()
	// TODO: compare table in each recursion
	if moves.Down {
		et.MoveDown()
	}
	if moves.Up {
		et.MoveUp()
	}
	if moves.Left {
		et.MoveLeft()
	}
	if moves.Right {
		et.MoveRight()
	}
}
