package main

import (
	"errors"
	"fmt"
)

//适用于没有删除的情况
//红黑树的增删查改最优
type AVLnode struct {
	//数据
	Data interface{}
	//左右指针
	Left  *AVLnode
	Right *AVLnode
	//高度
	height int
}

//函数指针类型 对比大小
type comparator func(a, b interface{}) int

//函数指针
var compare comparator

//比较大小
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func NewNode(data interface{}) *AVLnode {
	//新建节点
	node := new(AVLnode)
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.height = 1
	return node

}

//新建AVLtree
func NewAVLTree(data interface{}, myfunc comparator) (*AVLnode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("不能为空")
	}

	compare = myfunc
	return NewNode(data), nil

}

func (avlnode *AVLnode) Getall()[]interface{}{
	values := make([]interface{}, 0)
	return AddValues(values, avlnode)
}

func AddValues(values []interface{}, avlnode*AVLnode)[]interface{}{
	if avlnode != nil{
		values = AddValues(values, avlnode.Left)
		values = append(values, avlnode.Data)
		fmt.Println(avlnode.Data, avlnode.height)
		values = AddValues(values, avlnode.Right)

	}
	return values
}

//左旋 逆时针
func (avlnode *AVLnode) LeftRotate() *AVLnode {
	headnode := avlnode.Right
	avlnode.Right = headnode.Left
	headnode.Left = avlnode

	//跟新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode

}

//右旋 顺时针
func (avlnode *AVLnode) RightRotate() *AVLnode {
	//保存左边节点
	headnode := avlnode.Left
	avlnode.Left = headnode.Right
	headnode.Right = avlnode

	//跟新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode
}

//两次左旋

//两次右旋

//先左旋再右旋
func (avlnode *AVLnode) LeftThenRightRotate() *AVLnode {
	sonheadnode := avlnode.Left.LeftRotate()
	avlnode.Left = sonheadnode
	return avlnode.RightRotate()
}

//先右旋再左旋
func (avlnode *AVLnode) RightThenLeftRotate() *AVLnode {
	sonheadnode := avlnode.Right.RightRotate()
	avlnode.Right = sonheadnode
	return avlnode.LeftRotate()
}

//自动处理不平衡 差距为1 平衡 差距为2 不平衡
func (avlnode *AVLnode) adjust() *AVLnode {
	if avlnode.Right.GetHeight()-avlnode.Left.GetHeight() == 2 {
		if avlnode.Right.Right.GetHeight() > avlnode.Right.Left.GetHeight() {
			avlnode = avlnode.LeftRotate()
		} else {
			avlnode = avlnode.RightThenLeftRotate()
		}

	} else if avlnode.Left.GetHeight()-avlnode.Right.GetHeight() == 2 {
		if avlnode.Left.Left.GetHeight() > avlnode.Left.Right.GetHeight() {
			avlnode = avlnode.RightRotate()
		} else {
			avlnode = avlnode.LeftThenRightRotate()
		}
	}

	return avlnode
}

//数据插入
func (avlnode *AVLnode) Insert(value interface{}) *AVLnode {
	if avlnode == nil {
		newNode := &AVLnode{value, nil, nil, 1}
		return newNode
	}

	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Insert(value)
		avlnode = avlnode.adjust()
	case 1:
		avlnode.Right = avlnode.Right.Insert(value)
		avlnode = avlnode.adjust()
	case 0:
		fmt.Println("数据已经存在")
	}

	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight())+1
	return avlnode

}

//删除
func (avlnode *AVLnode)Delete(value interface{}) *AVLnode {
	if avlnode==nil{
		return nil
	}

	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Delete(value)
	case 1:
		avlnode.Right = avlnode.Right.Delete(value)
	case 0:
		// 删除在这里
		//左右都有节点
		if avlnode.Left != nil && avlnode.Right!=nil{
			avlnode.Data=avlnode.Right.FindMin().Data
			avlnode.Right = avlnode.Right.Delete(avlnode.Data)

		}else if avlnode.Left!=nil{
		//	左孩子存在 有孩子存在或不存在
			avlnode = avlnode.Left

		}else{
		//	只有一个右孩子
			avlnode = avlnode.Right

		}
	}

	if avlnode!=nil{
		avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight())+1
		//自动平衡
		avlnode = avlnode.adjust()
	}

	return avlnode

}


func (avlnode *AVLnode) Find(data interface{}) *AVLnode {
	var finded *AVLnode = nil
	switch compare(data, avlnode.Data) {
	case -1:
		finded = avlnode.Left.Find(data)
	case 1:
		finded = avlnode.Right.Find(data)
	case 0:
		return avlnode
	}
	return finded
}

func (avlnode *AVLnode) FindMin() *AVLnode {
	var finded *AVLnode
	if avlnode.Left != nil {
		//递归调用
		finded = avlnode.Left.FindMin()
	} else {
		finded = avlnode
	}
	return finded
}

func (avlnode *AVLnode) FindMax() *AVLnode {
	var finded *AVLnode
	if avlnode.Right != nil {
		//递归调用
		finded = avlnode.Right.FindMax()
	} else {
		finded = avlnode
	}
	return finded
}

//抓取数据
func (avlnode *AVLnode) Getdata() interface{} {
	if avlnode==nil{
		return 0
	}
	return avlnode.Data
}

//设置
func (avlnode *AVLnode) Setdata(data interface{}) {
	avlnode.Data = data
}

func (avlnode *AVLnode) GetLeft() *AVLnode {
	return avlnode.Left
}
func (avlnode *AVLnode) GetHeight() int {
	if avlnode==nil{
		return 0
	}
	return avlnode.height
}

func (avlnode *AVLnode) GetRight() *AVLnode {
	if avlnode==nil{
		return nil
	}
	return avlnode.Right
}
