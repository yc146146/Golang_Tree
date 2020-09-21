package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	key    interface{}
	value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

type SplayTree interface {
	//设定根节点
	SetRoot(n *Node)
	//返回根节点
	GetRoot() *Node
	//排序
	Ord(key1, key2 interface{}) int
}

//伸展树
type MySpalyTree struct {
	root *Node
}

func (ST *MySpalyTree) SetRoot(n *Node) {
	ST.root = n

}

func (ST *MySpalyTree) GetRoot() *Node {
	return ST.root

}
func (ST *MySpalyTree) Ord(key1, key2 interface{}) int {
	if key1.(int) < key2.(int) {
		return -1
	} else if key1.(int) == key2.(int) {
		return 0
	} else {
		return 1
	}
}

//搜索二叉树 返回节点
func Search(ST SplayTree, key interface{}) *Node {
	return SeachNode(ST, key, ST.GetRoot())
}

//Find 二叉树
func SeachNode(ST SplayTree, key interface{}, n *Node) *Node {
	if n == nil {
		return nil

	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return n
		case -1:
			//跳转到左边
			return SeachNode(ST, key, n.Left)
		case 1:
			return SeachNode(ST, key, n.Right)
		}
		return nil
	}
}

//查找实现返回数据
func Find(ST SplayTree, key interface{}) interface{} {
	return SeachNode(ST, key, ST.GetRoot())
}

//Find 二叉树
func FindNode(ST SplayTree, key interface{}, n *Node) interface{} {
	if n == nil {
		return nil

	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return n.value
		case -1:
			//跳转到左边
			return FindNode(ST, key, n.Left)
		case 1:
			return FindNode(ST, key, n.Right)
		}
		return nil
	}
}

func Insert(ST SplayTree, key interface{}, value interface{}) error {
	if Search(ST, key) != nil {
		return errors.New("要插入的已存在")
	}

	//调用插入
	n := InsertNode(ST, key, value, ST.GetRoot())
	//伸展

	fmt.Println(n)

	return nil

}

//插入数据
func InsertNode(ST SplayTree, key interface{}, value interface{}, n *Node) *Node {
	if n == nil {
		_n := new(Node)
		_n.key = key
		_n.value = value
		ST.SetRoot(_n)
		return ST.GetRoot()
	}

	switch ST.Ord(key, n.key) {
	case 0:
		//数据已经存在
		return nil
	case -1:
		//递归下去
		if n.Left == nil {
			n.Left = new(Node)
			n.Left.key = key
			n.Left.value = value
			//设定父亲节点
			n.Left.Parent = n
			return n.Left
		} else {
			//插入数据
			return InsertNode(ST, key, value, n.Left)
		}
	case 1:
		if n.Right == nil {
			n.Right = new(Node)
			n.Right.key = key
			n.Right.value = value
			//设定父亲节点
			n.Right.Parent = n
			return n.Right
		} else {
			return InsertNode(ST, key, value, n.Right)
		}
	}
	return nil
}

//删除数据
func Delete(ST SplayTree, key interface{}) error {
	n := Search(ST, key)
	if n == nil {
		return errors.New("要删除的不存在")
	} else {
		//保存父亲节点
		p := n.Parent
		if n.Left != nil {
			//寻找左边的最大值
			iop := InOrderPredecessor(n.Left)
			//交换节点
			Swap(n, iop)
			Remove(ST, iop)
		} else if n.Right != nil {
			//寻找右边最小值
			ios := InOrderSucccecessor(n.Right)
			Swap(n, ios)
			Remove(ST, ios)
		} else {
			Remove(ST, n)

		}
		if p != nil {
			//	伸展

		}
		return nil
	}

}

//删除节点
func Remove(ST SplayTree, n *Node) {
	var isRoot bool
	var isLeft bool
	//判断是否根节点
	isRoot = (n == ST.GetRoot())
	if isRoot != true {
		//判断是否左边节点
		isLeft = (n == n.Parent.Left)
	}
	if isRoot != true {
		if isLeft == true {
			if n.Left != nil {
				n.Parent.Left = n.Left
				n.Left.Parent = n.Parent

			} else if n.Right != nil {
				n.Parent.Left = n.Right
				n.Right.Parent = n.Parent

			} else {
				//叶子节点 左右都为空
				n.Parent.Left = nil

			}

		} else {
			if n.Left != nil {
				n.Parent.Right = n.Left
				n.Left.Parent = n.Parent
			} else if n.Right != nil {
				n.Parent.Right = n.Right
				n.Right.Parent = n.Parent
			} else {
				n.Parent.Right = nil
			}

		}
	}
	n = nil

}

//交换
func Swap(n1, n2 *Node) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

//取得最大
func InOrderPredecessor(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return InOrderPredecessor(n.Right)
	}
}

//取得最小
func InOrderSucccecessor(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return InOrderSucccecessor(n.Left)
	}
}

//伸展
func Splay(ST SplayTree, n *Node) {
	for n != ST.GetRoot() {
		if n.Parent == ST.GetRoot() && n.Parent.Left == n {
			ZigL(ST, n)
		} else if n.Parent == ST.GetRoot() && n.Parent.Right == n {
			ZigR(ST, n)
		} else if n.Parent.Left == n && n.Parent.Parent.Left == n.Parent {
			ZigZigL(ST, n)

		} else if n.Parent.Right == n && n.Parent.Parent.Right == n.Parent {
			ZigZigR(ST, n)
		} else if n.Parent.Right == n && n.Parent.Parent.Left == n.Parent {
			ZigZigLR(ST, n)
		} else {
			ZigZigRL(ST, n)
		}
	}
}

func ZigL(ST SplayTree, n *Node) {
	//存储左边数据  n旋转到根节点
	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}

	n.Parent.Parent = n
	n.Right = n.Parent
	n.Parent = nil

	ST.SetRoot(n)
}

func ZigR(ST SplayTree, n *Node) {

	//存储左边数据  n旋转到根节点
	n.Parent.Right = n.Left
	if n.Left != nil {
		n.Left.Parent = n.Parent
	}

	n.Parent.Parent = n
	n.Left = n.Parent
	n.Parent = nil

	ST.SetRoot(n)
}

func ZigZigL(ST SplayTree, n *Node) {
	//访问第三级节点
	gg := n.Parent.Parent.Parent
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.Left == n.Parent.Parent)
	}

	//备份left
	n.Parent.Parent.Left = n.Parent.Right
	if n.Parent.Right != nil {
		n.Parent.Right.Parent = n.Parent.Parent
	}

	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}

	n.Parent.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Right = n.Parent
	n.Parent.Parent = n
	n.Parent = gg

	//判断树，
	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.Left = n

	} else {
		gg.Right = n

	}
}

func ZigZigR(ST SplayTree, n *Node) {
	//访问第三级节点
	gg := n.Parent.Parent.Parent
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.Left == n.Parent.Parent)
	}

	//备份left
	n.Parent.Parent.Right = n.Parent.Left
	if n.Parent.Left != nil {
		n.Parent.Right.Parent = n.Parent.Parent
	}

	n.Parent.Right = n.Left
	if n.Left != nil {
		n.Left.Parent = n.Parent
	}

	n.Parent.Left = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Left = n.Parent
	n.Parent.Parent = n
	n.Parent = gg

	//判断树，
	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.Left = n

	} else {
		gg.Right = n

	}
}

func ZigZigLR(ST SplayTree, n *Node) {
	//访问第三级节点
	gg := n.Parent.Parent.Parent
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.Left == n.Parent.Parent)
	}

	//备份left
	n.Parent.Parent.Left = n.Parent.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent.Parent
	}

	n.Parent.Right = n.Left
	if n.Left != nil {
		n.Left.Parent = n.Parent
	}

	n.Left = n.Parent
	n.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg

	//判断树，
	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.Left = n

	} else {
		gg.Right = n

	}
}

func ZigZigRL(ST SplayTree, n *Node) {
	//访问第三级节点
	gg := n.Parent.Parent.Parent
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = (gg.Left == n.Parent.Parent)
	}

	//备份left
	n.Parent.Parent.Right = n.Parent.Left
	if n.Left != nil {
		n.Left.Parent = n.Parent.Parent
	}

	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}

	n.Right = n.Parent
	n.Left = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg

	//判断树，
	if isRoot == true {
		ST.SetRoot(n)
	} else if isLeft == true {
		gg.Left = n

	} else {
		gg.Right = n

	}
}

//显示树
func Print(ST SplayTree) {
	PrintNode(ST.GetRoot(), 0)
}

//显示节点
func PrintNode(n *Node, level int) {
	if n == nil {
		return
	}

	PrintNode(n.Left, level+1)
	//打印数据
	fmt.Println(strings.Repeat("-", 2*level), n.key, n.value)
	PrintNode(n.Right, level+1)
}

func main() {
	ST := new(MySpalyTree)
	for i := 0; i < 36; i++ {
		err := Insert(ST, rand.Int()%100, "hello word")
		if err != nil {
			fmt.Println(err)
		} else {
			Print(ST)
		}
	}
	for i := 0; i < 36; i++ {
		err := Delete(ST, i)
		if err != nil {
			fmt.Println(err)
		} else {
			Print(ST)
		}
	}
}
