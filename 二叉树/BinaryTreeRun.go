package main

import "fmt"

func main() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil,nil}
	node2 := &Node{2, nil,nil}
	node3 := &Node{6, nil,nil}
	node4 := &Node{1, nil,nil}
	node5 := &Node{3, nil,nil}
	node6 := &Node{5, nil,nil}
	node7 := &Node{7, nil,nil}
	bst.Root = node1
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	bst.Size = 7

	nodelast := bst.FindlowerstAncestor(bst.Root, node3, node4)
	fmt.Println(nodelast)

	fmt.Println((bst.GetDepth(bst.Root)))
}


func main2() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil,nil}
	node2 := &Node{2, nil,nil}
	node3 := &Node{6, nil,nil}
	node4 := &Node{1, nil,nil}
	node5 := &Node{3, nil,nil}
	node6 := &Node{5, nil,nil}
	node7 := &Node{7, nil,nil}
	bst.Root = node1
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	bst.Size = 7


	//fmt.Println("-------------------------")
	//bst.InOrder()
	//fmt.Println("-------------------------")
	//fmt.Println(bst.InOrderNoRecursion())

	//fmt.Println("-------------------------")
	//bst.PreOrder()
	//fmt.Println("-------------------------")
	//fmt.Println(bst.PreOrderNoRecursion())
	//fmt.Println("-------------------------")
	//bst.PostOrder()
	//fmt.Println("-------------------------")
	//fmt.Println(bst.PostOrderNoRecursion())
	fmt.Println("-------------------------")
	//fmt.Println(bst.String())
	//bst.Levelshow()
	bst.Stackshow(bst.Root)
}



func main11() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil,nil}
	node2 := &Node{2, nil,nil}
	node3 := &Node{6, nil,nil}
	node4 := &Node{1, nil,nil}
	node5 := &Node{3, nil,nil}
	node6 := &Node{5, nil,nil}
	node7 := &Node{7, nil,nil}
	bst.Root = node1
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	bst.Size = 7

	//for i:=1;i<=7;i++{
	//	bst.Add(i)
	//}


	//bst.Add(4)
	//bst.Add(6)
	//bst.Add(5)
	//bst.Add(7)
	//
	//bst.Add(2)
	//bst.Add(1)
	//bst.Add(3)


	fmt.Println(bst.FindMin())
	fmt.Println(bst.FindMax())
	fmt.Println(bst.Isin(3))
	fmt.Println(bst.Isin(31))

	//fmt.Println(bst.RemoveMin())
	//fmt.Println(bst.RemoveMax())

	//fmt.Println()
	bst.Remove(4)

	fmt.Println("-------------------------")
	bst.InOrder()
	//fmt.Println("-------------------------")
	//bst.PreOrder()
	//fmt.Println("-------------------------")
	//bst.PostOrder()
	fmt.Println("-------------------------")
	fmt.Println(bst.String())
	fmt.Println("-------------------------")

}
