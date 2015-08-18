package main

import (
	"errors"
)

// EightTable stores numbers in 8-puzzle
type EightTable struct {
	table [3][3]int
}
type Position struct {
	Row    int
	Column int
}

const (
	TABLE_SIZE = 3
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
