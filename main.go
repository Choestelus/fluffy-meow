package main

import "fmt"

func main() {
	et := NewTable()
	fmt.Println(et)
	fmt.Println("blank Position = ", et.FindBlank())
}
