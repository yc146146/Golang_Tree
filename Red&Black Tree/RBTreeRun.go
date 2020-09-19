package main

import "fmt"

func main() {
	rbtree := NewRBTree()
	for i:=0;i<100000;i++{
		rbtree.Insert(Int(i))
	}

	for i:=0;i<90000;i++{
		rbtree.Delete(Int(i))
	}

	fmt.Println(rbtree.GetDepth())
}