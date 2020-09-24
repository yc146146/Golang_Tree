package main

import "sort"

//存储数据
type kc struct {
	//数据
	key int
	//接口类型
	child node
}

//数组
type kcs [MaxKC + 1]kc

func (kvs *kcs) Len() int {
	return len(kvs)
}

//判断大小
func (kvs *kcs) Less(i, j int) bool {

	//中间节点 数组第一个空着
	if kvs[i].key == 0 {
		return false
	}

	if kvs[j].key == 0 {
		return true
	}

	return kvs[i].key < kvs[j].key
}

//中间节点数据结构
type interiorNode struct {
	//数组存储数据
	kcs kcs
	//存储元素数量
	count int
	//指向父亲节点
	parent *interiorNode
}

//新建一个中间节点
func NewinteriorNode(p *interiorNode, largestChild node) *interiorNode {
	in := &interiorNode{parent: p, count: 1}
	if largestChild != nil {
		//设置插入节点  插入叶子节点
		in.kcs[0].child = largestChild
	}
	return in
}

func (in *interiorNode) find(key int) (int, bool) {
	//是个函数主要进行数据对比
	myfunc := func(i int) bool {
		return in.kcs[i].key >= key
	}

	//实现查询
	i := sort.Search(in.count-1, myfunc)

	return i, true
}

//插入节点
func (in *interiorNode) insert(key int, child node) (int, *interiorNode, bool) {
	//确定位置
	i, _ := in.find(key)
	if !in.full() {
		//	数据插入之前要整体向后移动
		copy(in.kcs[i+1:], in.kcs[i:in.count])
		//设置子节点分裂以后的元素 设置key
		in.kcs[i].key = key
		in.kcs[i].child = child
		//设定父亲节点
		child.SetParent(in)
		in.count++
		return 0, nil, false
	} else {
		//存储到最后
		in.kcs[MaxKC].key = key
		in.kcs[MaxKC].child = child

		child.SetParent(in)
		next, midkey := in.split()
		//返回中间节点
		return midkey, next, true

	}

}

func (in *interiorNode) split() (*interiorNode, int) {
	//	节点分裂 节点插入正确位置
	//	确保有序
	sort.Sort(&in.kcs)

	//	取得中间节点
	midIndex := MaxKC / 2
	//取得中间节点
	midChild := in.kcs[midIndex].child
	//取得键值
	midkey := in.kcs[midIndex].key
	//	新建一个中间节点

	next := NewinteriorNode(nil, nil)
	//拷贝数据
	copy(next.kcs[0:], in.kcs[midIndex+1:])
	//数据的初始化
	in.InitArray(midIndex + 1)
	//下一个节点数量
	next.count = MaxKC - midIndex

	//新开辟节点的每个叶子节点的节点祖先设置为next
	for i := 0; i < next.count; i++ {
		next.kcs[i].child.SetParent(next)
	}
	in.count = midIndex + 1
	//设置为0 预留一个
	in.kcs[in.count-1].key = 0
	in.kcs[in.count-1].child = midChild
	midChild.SetParent(in)
	return next, midkey
}

func (kvs *kcs) Swap(i, j int) {
	kvs[i], kvs[j] = kvs[j], kvs[i]
}

func (in *interiorNode) full() bool {
	return in.count == MaxKC
}

func (in *interiorNode) Parent() *interiorNode {
	return in.parent
}

func (in *interiorNode) SetParent(p *interiorNode) {
	in.parent = p
}

func (in *interiorNode) CountNum() int {
	return in.count
}

//初始化数组 数组每个元素初始化
func (in *interiorNode) InitArray(num int) {
	for i := num; i < len(in.kcs); i++ {
		in.kcs[i] = kc{}
	}
}
