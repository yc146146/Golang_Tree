package main

import (
	"fmt"
	"math/rand"
	"time"
)

//B树的节点
type BtreeNode struct {
	Leaf     bool
	N        int
	keys     []int
	Children []*BtreeNode
}

func NewBtreeNode(n int, branch int, leaf bool) *BtreeNode {

	return &BtreeNode{leaf,
		n,
		make([]int, branch*2-1),
		make([]*BtreeNode, branch*2)}
}

//搜索b树的节点
func (btreenode *BtreeNode) Search(key int) (mynode *BtreeNode, idx int) {
	i := 0
	//找到合适的位置,找到最后一个小于key的 i之后的就是大于等于
	for i < btreenode.N && btreenode.keys[i] < key {
		i += 1
	}

	if i < btreenode.N && btreenode.keys[i] == key {
		//找到了
		mynode, idx = btreenode, i

	} else if btreenode.Leaf == false {
		//	进入孩子叶子继续搜索
		mynode, idx = btreenode.Children[i].Search(key)
	}
	return
}

func (parent *BtreeNode) Split(branch int, idx int) {

	//孩子节点
	full := parent.Children[idx]
	//新建一个节点
	newnode := NewBtreeNode(branch-1, branch, full.Leaf)
	for i := 0; i < branch-1; i++ {
		//数据的移动,跳过一个分支大小
		newnode.keys[i] = full.keys[i+branch]
		newnode.Children[i] = full.Children[i+branch]
	}
	newnode.Children[branch-1] = full.Children[branch*2-1]
	full.N = branch - 1
	//新增一个key到children
	for i := parent.N; i > idx; i-- {
		parent.Children[i] = parent.Children[i-1]
		//从后往前移动
		parent.keys[i+1] = parent.keys[i]
	}
	parent.keys[idx] = full.keys[branch-1]
	//插入数据,增加总量
	parent.Children[idx+1] = newnode
	parent.N++
}

//节点插入数据
func (btreenode *BtreeNode) InsertNonFull(branch int, key int) {
	if btreenode == nil{
		return
	}
	//记录叶子节点总量
	i := btreenode.N
	//是叶子或者不是叶子
	if btreenode.Leaf {
		for i > 0 && key < btreenode.keys[i-1] {
			//从后往前移动
			btreenode.keys[i] = btreenode.keys[i-1]
			//从后往前移动
			i--
		}
		//插入数量
		btreenode.keys[i] = key
		//总量+1
		btreenode.N++

	} else {
		for i > 0 && key < btreenode.keys[i-1] {
			i--
		}
		//找到下标
		c := btreenode.Children[i]
		if c!=nil && c.N == 2*branch-1 {
			btreenode.Split(branch, i)
			if key > btreenode.keys[i] {
				i++
			}
		}
		//递归插入孩子叶子
		btreenode.Children[i].InsertNonFull(branch, key)
	}

}

//节点显示成字符串
func (btreeNode *BtreeNode)String()string{
	return fmt.Sprintf("n=%d,leaf=%v,Children=%v \n",btreeNode.N,btreeNode.keys,btreeNode.Children )

}

//b树
type Btree struct {
	//根节点
	Root *BtreeNode
	//分支数量
	branch int
}

//插入
func (tree *Btree) Insert(key int) {
	root := tree.Root
	if root.N == 2*tree.branch-1 {
		s := NewBtreeNode(0, tree.branch, false)
		tree.Root = s
		s.Children[0] = root
		//拆分整合
		s.Split(tree.branch, 0)
		root.InsertNonFull(tree.branch, key)
	} else {
		root.InsertNonFull(tree.branch, key)
	}
}

//查找
func (tree *Btree) Search(key int) (n *BtreeNode, idx int) {
	return tree.Root.Search(key)
}

//返回字符串
func (tree *Btree) String() string {
	//返回树的字符串
	return tree.Root.String()
}

//新建b树
func NewBtree(branch int) * Btree{
	return &Btree{NewBtreeNode(0,branch,true), branch}
}

func main() {
	mybtree := NewBtree(100000)

	for i:=100000;i>0;i--{
		mybtree.Insert(rand.Int()%100000)
	}
	fmt.Println(mybtree.String())



	for i:=0;i<10000;i++{
		startTime := time.Now()
		fmt.Println(mybtree.Search(i))
		fmt.Println("公用", time.Since(startTime))
	}

}
