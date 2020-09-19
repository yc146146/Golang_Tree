package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	//数据
	Data int
	//指向左边节点
	Left *Node
	//指向右边节点
	Right *Node
}

type BinaryTree struct {
	//根节点
	Root *Node
	//	数据的数量
	Size int
}

//新建一个二叉树
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{}
	bst.Size = 0
	bst.Root = nil
	return bst

}

//获取二叉树的大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

//判断是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

//跟节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

//插入节点
func (bst *BinaryTree) add(n *Node, data int) *Node {
	if n == nil {
		bst.Size++
		//return &Node{Data:data}
		return &Node{data, nil, nil}
	} else {
		if data < n.Data {
			//比我小左边
			n.Left = bst.add(n.Left, data)
		} else if data > n.Data {
			n.Right = bst.add(n.Right, data)
		}

		return n
	}
}

//判断是否存在
func (bst *BinaryTree) Isin(data int) bool {
	//从根节点查找
	return bst.isin(bst.Root, data)
}

//判断是否存在
func (bst *BinaryTree) isin(n *Node, data int) bool {
	if n == nil {
		//树是空树
		return false
	}
	if data == n.Data {
		return true
	} else if data < n.Data {
		return bst.isin(n.Left, data)
	} else {
		return bst.isin(n.Right, data)
	}
}

//判断是否存在
func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findMax(bst.Root).Data
}

func (bst *BinaryTree) findMax(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findMax(n.Right)
	}

}

func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findMin(bst.Root).Data
}

func (bst *BinaryTree) findMin(n *Node) *Node {
	if n.Left == nil {
		return n
	} else {
		return bst.findMin(n.Left)
	}
}

//前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preOrder(bst.Root)
}

func (bst *BinaryTree) PreOrderNoRecursion() []int {
	mybst := bst.Root
	mystack := list.New() //生成一个栈
	//生成数组 容纳中序数据
	res := make([]int, 0)
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			res = append(res, mybst.Data)
			mystack.PushBack(mybst)
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()
			//实例化
			mybst = v.Value.(*Node)

			//追加
			mybst = mybst.Right
			mystack.Remove(v) //删除
		}
	}

	return res

}

func (bst *BinaryTree) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	bst.preOrder(node.Left)
	bst.preOrder(node.Right)

}

//中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inOrder(bst.Root)
}

func (bst *BinaryTree) InOrderNoRecursion() []int {
	mybst := bst.Root
	mystack := list.New() //生成一个栈
	//生成数组 容纳中序数据
	res := make([]int, 0)
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst)
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()
			//实例化
			mybst = v.Value.(*Node)
			res = append(res, mybst.Data)
			//追加
			mybst = mybst.Right
			mystack.Remove(v) //删除
		}
	}

	return res

}
func (bst *BinaryTree) inOrder(node *Node) {
	if node == nil {
		return
	}

	bst.inOrder(node.Left)
	fmt.Println(node.Data)
	bst.inOrder(node.Right)
}

//后续遍历
func (bst *BinaryTree) PostOrder() {
	bst.postOrder(bst.Root)
}
func (bst *BinaryTree) PostOrderNoRecursion() []int {
	mybst := bst.Root
	mystack := list.New() //生成一个栈
	//生成数组 容纳中序数据
	res := make([]int, 0)
	var PreVisited *Node //提前访问的节点

	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst)
			mybst = mybst.Left
		}
		//取出战中节点
		v := mystack.Back()
		top := v.Value.(*Node)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && PreVisited == top.Left) || (PreVisited == top.Right) {
			res = append(res, top.Data)
			PreVisited = top
			mystack.Remove(v)
		} else {
			mybst = top.Right
		}
	}

	return res

}

func (bst *BinaryTree) postOrder(node *Node) {

	if node == nil {
		return
	}

	bst.postOrder(node.Left)
	bst.postOrder(node.Right)
	fmt.Println(node.Data)

}

func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer
	//调用函数实现遍历
	bst.GenerateBSTstring(bst.Root, 0, &buffer)
	return buffer.String()
}

func (bst *BinaryTree) GenerateBSTstring(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		//空节点
		//buffer.WriteString(bst.GenerateDepthstring(depth)+"nil\n")
		return
	}
	//写入字符串 保存树的深度
	bst.GenerateBSTstring(node.Left, depth+1, buffer)
	buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Data) + "\n")
	bst.GenerateBSTstring(node.Right, depth+1, buffer)
}

func (bst *BinaryTree) GenerateDepthstring(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}

	return buffer.String()
}

func (bst *BinaryTree) RemoveMin() int {
	ret := bst.FindMin()
	bst.Root = bst.removemin(bst.Root)
	return ret

}

//删除最小
func (bst *BinaryTree) removemin(n *Node) *Node {
	if n.Left == nil {
		//	删除
		//备份右边的节点
		rightNode := n.Right
		bst.Size--

		return rightNode
	}
	n.Left = bst.removemin(n.Left)
	return n

}

func (bst *BinaryTree) RemoveMax() int {
	ret := bst.FindMax()
	bst.Root = bst.removemax(bst.Root)
	return ret

}

func (bst *BinaryTree) removemax(n *Node) *Node {
	if n.Right == nil {
		//	删除
		//备份右边的节点
		leftNode := n.Left
		bst.Size--

		return leftNode
	}
	n.Right = bst.removemax(n.Right)
	return n
}

func (bst *BinaryTree) Remove(data int) {
	bst.Root = bst.remove(bst.Root, data)

}

func (bst *BinaryTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil
	}

	if data < n.Data {
		n.Left = bst.remove(n.Left, data)
		return n
	} else if data > n.Data {
		n.Right = bst.remove(n.Right, data)
		return n
	} else {
		//处理左边为空
		if n.Left == nil {
			rightNode := n.Right
			n.Right = nil

			bst.Size--

			return rightNode
		}

		//处理右边为空
		if n.Right == nil {
			leftNode := n.Left
			n.Left = nil
			bst.Size--

			return leftNode
		}

		//	左右节点都不为空
		//找出比我小的节点
		oknode := bst.findMin(n.Right)
		oknode.Right = bst.removemin(n.Right)
		//删除
		oknode.Left = n.Left

		n.Left = nil
		n.Right = nil
		return oknode
	}

}

func (bst *BinaryTree) Levelshow() {
	bst.levelshow(bst.Root)
}

func (bst *BinaryTree) levelshow(n *Node) {
	myqueue := list.New()
	myqueue.PushBack(n)
	for myqueue.Len() > 0 {
		//前面取出数据
		left := myqueue.Front()
		right := left.Value
		myqueue.Remove(left)
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data)
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) Stackshow(n *Node) {
	myqueue := list.New()
	myqueue.PushBack(n)
	for myqueue.Len() > 0 {
		//前面取出数据
		//此时是栈
		left := myqueue.Back()
		right := left.Value
		myqueue.Remove(left)
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data)
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) FindlowerstAncestor(root *Node, nodea *Node, nodeb *Node) *Node {
	if root == nil {
		return nil
	}
	if root == nodea || root == nodeb {
		//两者有一个节点是根节点
		return root
	}

	left := bst.FindlowerstAncestor(root.Left, nodea, nodeb)
	right := bst.FindlowerstAncestor(root.Right, nodea, nodeb)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}

}

func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	lengthleft := bst.GetDepth(root.Left)
	lengthright := bst.GetDepth(root.Right)

	if lengthleft > lengthright {
		return lengthleft + 1
	} else {
		return lengthright + 1
	}
}
