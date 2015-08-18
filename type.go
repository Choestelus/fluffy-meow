package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// EightTable stores numbers in 8-puzzle
type EightTable struct {
	table [3][3]int
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
	for i, e := range et.table {
		for ii := range e {
			et.table[i][ii] = 3*i + ii
		}
	}
	return et
}

func (et EightTable) Print() {
	for _, e := range et.table {
		fmt.Println("-->", e)
	}
}

func (et EightTable) FindBlank() Position {
	var pos Position
	for i := range et.table {
		for j := range et.table[i] {
			if et.table[i][j] == 0 {
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
			}
		}
		if r == MOVE_DOWN {
			err := et.MoveDown()
			if err != nil {
			} else {
				successCount++
			}
		}
		if r == MOVE_LEFT {
			err := et.MoveLeft()
			if err != nil {
			} else {
				successCount++
			}
		}
		if r == MOVE_RIGHT {
			err := et.MoveRight()
			if err != nil {
			} else {
				successCount++
			}
		}

	}
}

func (et *EightTable) MoveLeft() error {
	bpos := et.FindBlank()
	if bpos.Column > 0 {
		et.table[bpos.Row][bpos.Column], et.table[bpos.Row][bpos.Column-1] = et.table[bpos.Row][bpos.Column-1], et.table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to left")
	}
	return nil
}
func (et *EightTable) MoveRight() error {
	bpos := et.FindBlank()
	if bpos.Column < TABLE_SIZE-1 {
		et.table[bpos.Row][bpos.Column], et.table[bpos.Row][bpos.Column+1] = et.table[bpos.Row][bpos.Column+1], et.table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to right")
	}
	return nil
}
func (et *EightTable) MoveUp() error {
	bpos := et.FindBlank()
	if bpos.Row > 0 {
		et.table[bpos.Row][bpos.Column], et.table[bpos.Row-1][bpos.Column] = et.table[bpos.Row-1][bpos.Column], et.table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to up")
	}
	return nil
}
func (et *EightTable) MoveDown() error {
	bpos := et.FindBlank()
	if bpos.Row < TABLE_SIZE-1 {
		et.table[bpos.Row][bpos.Column], et.table[bpos.Row+1][bpos.Column] = et.table[bpos.Row+1][bpos.Column], et.table[bpos.Row][bpos.Column]
	} else {
		return errors.New("cannot move blank space to up")
	}
	return nil
}
