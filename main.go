package main

import "fmt"

func main() {
	et := NewTable()
	err := et.MoveLeft()
	if err != nil {
		fmt.Println("error:", err)
	}
	et.Print()
}
