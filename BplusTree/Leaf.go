package main

import "sort"

type kv struct {
	key   int
	value string
}

type kvs [MaxKC]kv

func (kvs *kvs) Len() int {
	return len(kvs)
}

//交换数据
func (kvs *kvs) Swap(i, j int) {
	kvs[i], kvs[j] = kvs[j], kvs[i]
}

//判断大小
func (kvs *kvs) Less(i, j int) bool {
	return kvs[i].key < kvs[j].key
}

//叶子节点
type LeafNode struct {
	//数据
	kvs kvs
	//数量
	count int
	//下一个节点
	next *LeafNode
	//父亲节点
	parent *interiorNode
}

//创建叶子节点
func NewLeafNode(parent *interiorNode) *LeafNode {
	return &LeafNode{parent: parent}
}

//查找数据
func (l *LeafNode) find(key int) (int, bool) {
	//是个函数主要进行数据对比
	myfunc := func(i int) bool {
		return l.kvs[i].key >= key
	}

	//实现查询
	i := sort.Search(l.count, myfunc)
	if i < l.count && l.kvs[i].key == key {
		return i, true
	}
	return i, false

}

//插入叶子节点
func (l *LeafNode) insert(key int, value string) (int, *LeafNode, bool) {
	i, ok := l.find(key)
	if ok {
		//	key数据已经存在 跟新value
		l.kvs[i].value = value
		return 0, nil, false
	}
	//叶子节点是否满了
	if !l.full(){
		//数组删除需要整体向后移动
		copy(l.kvs[i+1:], l.kvs[i:l.count])
		l.kvs[i].key = key
		l.kvs[i].value = value
		l.count++
		return 0, nil, false
	}else{
		//分裂叶子节点
		next := l.split()
		if key < next.kvs[0].key{
			l.insert(key,value)
		}else{
			next.insert(key,value)
		}
		return next.kvs[0].key,next,true
	}
}

func (l *LeafNode) full()bool{
	return l.count==MaxKV
}

func (l *LeafNode) Parent()*interiorNode{
	return l.parent
}

func (l *LeafNode) SetParent(p*interiorNode){
	l.parent = p
}

func (l *LeafNode) CountNum()int{
	return l.count
}

//初始化数组 数组每个元素初始化
func (l *LeafNode) InitArray(num int) {
	for i := num; i < len(l.kvs); i++ {
		l.kvs[i] = kv{}
	}
}

//叶子节点 分裂  123 4 567
func (l *LeafNode) split() *LeafNode {
	//新建一个右边节点
	next := NewLeafNode(nil)
	//复制数据到右边节点
	copy(next.kvs[0:], l.kvs[l.count/2+1:])
	//后半部数据清空
	l.InitArray(l.count/2 + 1)
	//下一个节点数量
	next.count = MaxKV - l.count/2 - 1
	//调整指针的指向
	next.next = l.next
	l.count = l.count/2 + 1
	l.next = next
	return next

}


