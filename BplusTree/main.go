package main

import "fmt"

func main() {
	bpt := NewBplusTree()
	for i:=0;i<10000;i++{
		bpt.Insert(i, "x")
	}
	fmt.Println(bpt.Count())
	fmt.Println(bpt.Search(3))
}
