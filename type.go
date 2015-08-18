package main

// EightTable stores numbers in 8-puzzle
type EightTable struct {
	table [3][3]int
}
type Position struct {
	Row    int
	Column int
}

const (
	MOVE_UP    = 0
	MOVE_DOWN  = 1
	MOVE_LEFT  = 2
	MOVE_RIGHT = 3
)

type Movables interface {
}

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

func (et *EightTable) FindBlank() Position {
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
func (et *EightTable) MoveLeft() {
}
func (et *EightTable) MoveRight() {
}
func (et *EightTable) MoveUp() {
}
func (et *EightTable) MoveDown() {
}
func (et *EightTable) Movables() {
}
