package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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
func (et *EightTable) Shuffle(n int) {
	var successCount int
	for i := 0; successCount < n; i++ {
		r := rand.Intn(4)
		if r == MOVE_UP {
			err := et.MoveUp()
			if err != nil {
			} else {
				successCount++
				fmt.Println("moved up")
			}
		}
		if r == MOVE_DOWN {
			err := et.MoveDown()
			if err != nil {
			} else {
				successCount++
				fmt.Println("moved down")
			}
		}
		if r == MOVE_LEFT {
			err := et.MoveLeft()
			if err != nil {
			} else {
				successCount++
				fmt.Println("moved left")
			}
		}
		if r == MOVE_RIGHT {
			err := et.MoveRight()
			if err != nil {
			} else {
				successCount++
				fmt.Println("moved right")
			}
		}

	}
}
func (et EightTable) Self() EightTable {
	return et
}

func (et *EightTable) MoveLeft() error {
	bpos := et.FindBlank()
	if bpos.Column > 0 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row][bpos.Column-1] = et.Table[bpos.Row][bpos.Column-1], et.Table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to left")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Left: true})
	return nil
}
func (et *EightTable) MoveRight() error {
	bpos := et.FindBlank()
	if bpos.Column < TABLE_SIZE-1 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row][bpos.Column+1] = et.Table[bpos.Row][bpos.Column+1], et.Table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to right")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Right: true})
	return nil
}

func (et *EightTable) MoveUp() error {
	bpos := et.FindBlank()
	if bpos.Row > 0 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row-1][bpos.Column] = et.Table[bpos.Row-1][bpos.Column], et.Table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to up" + strconv.Itoa(bpos.Row))
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Up: true})
	return nil
}

func (et *EightTable) MoveDown() error {
	bpos := et.FindBlank()
	if bpos.Row < TABLE_SIZE-1 {
		et.Table[bpos.Row][bpos.Column], et.Table[bpos.Row+1][bpos.Column] = et.Table[bpos.Row+1][bpos.Column], et.Table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to up")
	}
	et.MoveHistory = append(et.MoveHistory, Moves{Down: true})
	return nil
}
