package main

const (
	//叶子节点的最大存储数量 2^-1
	MaxKV = 255
	//中间节点最大存储数量
	MaxKC = 511
)


//接口设计
type node interface {
	//查找key
	find(key int)(int, bool)
	//返回父亲节点
	Parent() * interiorNode
	SetParent(*interiorNode)
	//判断是否满了
	full()bool
	//统计元素数量
	CountNum()int

}
