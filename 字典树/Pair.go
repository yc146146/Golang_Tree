package main

import (
	"fmt"
	"sync"
)

type KeyWordTreeNode struct {
	KeyWordIDs map[int64]bool
	//1,2,3,4
	Char string
	//父结点
	ParentKeyWordTreeNode *KeyWordTreeNode
	//子节点的集合
	SubKeyWordTreeNodes map[string]*KeyWordTreeNode
}

//初始化节点
func NewKeyWordTreeNode() *KeyWordTreeNode {
	return &KeyWordTreeNode{make(map[int64]bool, 0),
		"", nil, make(map[string]*KeyWordTreeNode, 0)}
}

//初始化节点设置内容以及父亲节点
func NewKeyWordTreeNodeWithParams(ch string, parent *KeyWordTreeNode) *KeyWordTreeNode {
	return &KeyWordTreeNode{make(map[int64]bool, 0),
		ch, parent, make(map[string]*KeyWordTreeNode, 0)}
}

type KeyWordTree struct {
	//根节点
	root *KeyWordTreeNode
	//映射关系
	kv KeyWordKV
	//开始映射
	char_begin_kv CharBeginKV
	//保护线程安全
	rw *sync.RWMutex
}

func NewKeyWordTree() *KeyWordTree {
	return &KeyWordTree{NewKeyWordTreeNode(),
		KeyWordKV{},
		CharBeginKV{},
		new(sync.RWMutex),
	}
}

func (stree *KeyWordTree) DebugOut() {
	//输出节点
	fmt.Println("s.kv", stree.kv)
	temproot := stree.root
	dfs(temproot)
}

//遍历文件树
func dfs(root *KeyWordTreeNode) {
	if root == nil {
		return
	} else {
		fmt.Println("s.root=", root.Char)
		fmt.Println("s.KeywordIds=", root.KeyWordIDs)
		for _, v := range root.SubKeyWordTreeNodes {
			dfs(v)
		}
	}
}
