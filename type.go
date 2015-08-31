package main

import (
	"fmt"
	"math/rand"
)

// EightTable stores numbers in 8-puzzle
type EightTable struct {
	Table       [3][3]int
	MoveHistory []Moves
}
type Position struct {
	Row    int
	Column int
}
type Moves struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

const (
	TABLE_SIZE = 3
	MOVE_UP    = 0
	MOVE_DOWN  = 1
	MOVE_LEFT  = 2
	MOVE_RIGHT = 3
)

// NewTable initialize puzzle table
func NewTable() EightTable {
	var et EightTable
	for i, e := range et.Table {
		for ii := range e {
			et.Table[i][ii] = 3*i + ii + 1
			if 3*i+ii+1 == 9 {
				et.Table[i][ii] = 0
			}
		}
	}
	return et
}

func (et EightTable) Print() {
	for _, e := range et.Table {
		fmt.Println("-->", e)
	}
}

func (m Moves) Print() string {
	str := ""
	if m.Up {
		str += "Up"
	}
	if m.Down {
		str += "Down"
	}

	if m.Left {
		str += "Left"
	}
	if m.Right {
		str += "Right"
	}
	return "[" + str + "]"
}

func (et EightTable) FindBlank() Position {
	var pos Position
	for i := range et.Table {
		for j := range et.Table[i] {
			if et.Table[i][j] == 0 {
				pos = Position{i, j}
			}
		}
	}
	return pos
}

func (et EightTable) Movables() Moves {
	pos := et.FindBlank()
	var moves Moves
	if pos.Column > 0 {
		moves.Left = true
	}
	if pos.Column < TABLE_SIZE-1 {
		moves.Right = true
	}
	if pos.Row > 0 {
		moves.Up = true
	}
	if pos.Row < TABLE_SIZE-1 {
		moves.Down = true
	}
	return moves
}
func (et EightTable) Shuffle(n int) (newEt EightTable) {
	var successCount int
	for i := 0; successCount < n; i++ {
		r := rand.Intn(4)
		canMove := et.Movables()
		if r == MOVE_UP && canMove.Up {
			et = et.MoveUp()
			successCount++
			// et.Print()
		}
		if r == MOVE_DOWN && canMove.Down {
			et = et.MoveDown()
			successCount++
			// et.Print()
		}
		if r == MOVE_LEFT && canMove.Left {
			et = et.MoveLeft()
			successCount++
			// et.Print()
		}
		if r == MOVE_RIGHT && canMove.Right {
			et = et.MoveRight()
			successCount++
			// et.Print()
		}

	}
	newEt = et
	return
}
func (et EightTable) Self() EightTable {
	return et
}
func (eet EightTable) MoveLeft() (et EightTable) {
	for i := 0; i < TABLE_SIZE; i++ {
		for j := 0; j < TABLE_SIZE; j++ {
			et.Table[i][j] = eet.Table[i][j]
		}
	}

	for _, e := range eet.MoveHistory {
		et.MoveHistory = append(et.MoveHistory, e)
	}
	bpos := et.FindBlank()
	if bpos.Column > 0 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row][bpos.Column-1] = et.Table[bpos.Row][bpos.Column-1], et.Table[bpos.Row][bpos.Column]
	} else {
		panic("Error Moving Left")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Left: true})
	return et
}

func (eet EightTable) MoveRight() (et EightTable) {
	for i := 0; i < TABLE_SIZE; i++ {
		for j := 0; j < TABLE_SIZE; j++ {
			et.Table[i][j] = eet.Table[i][j]
		}
	}
	for _, e := range eet.MoveHistory {
		et.MoveHistory = append(et.MoveHistory, e)
	}
	bpos := et.FindBlank()
	if bpos.Column < TABLE_SIZE-1 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row][bpos.Column+1] = et.Table[bpos.Row][bpos.Column+1], et.Table[bpos.Row][bpos.Column]
	} else {
		panic("Error Moving Right")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Right: true})
	return et
}

func (eet EightTable) MoveUp() (et EightTable) {
	for i := 0; i < TABLE_SIZE; i++ {
		for j := 0; j < TABLE_SIZE; j++ {
			et.Table[i][j] = eet.Table[i][j]
		}
	}
	for _, e := range eet.MoveHistory {
		et.MoveHistory = append(et.MoveHistory, e)
	}
	bpos := et.FindBlank()
	if bpos.Row > 0 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row-1][bpos.Column] = et.Table[bpos.Row-1][bpos.Column], et.Table[bpos.Row][bpos.Column]
	} else {
		panic("Error Moving Up")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Up: true})
	return et
}

func (eet EightTable) MoveDown() (et EightTable) {
	for i := 0; i < TABLE_SIZE; i++ {
		for j := 0; j < TABLE_SIZE; j++ {
			et.Table[i][j] = eet.Table[i][j]
		}
	}
	for _, e := range eet.MoveHistory {
		et.MoveHistory = append(et.MoveHistory, e)
	}
	bpos := et.FindBlank()
	if bpos.Row < TABLE_SIZE-1 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row+1][bpos.Column] = et.Table[bpos.Row+1][bpos.Column], et.Table[bpos.Row][bpos.Column]
	} else {
		panic("error moving down")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Down: true})
	return et
}
