package main

import "fmt"

func main() {
	et := NewTable()
	err := et.MoveLeft()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(et)
	fmt.Println("blank Position = ", et.FindBlank())
}
