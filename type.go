package main

// EightTable stores numbers in 8-puzzle
type EightTable struct {
	table [3][3]int
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
