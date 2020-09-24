package main

//定义
type BPlustree map[int]node



func NewBplusTree()*BPlustree{
	//初始化
	bt := BPlustree{}
	//叶子节点
	leaf := NewLeafNode(nil)
	//中间节点
	r := NewinteriorNode(nil, leaf)
	//设定父亲节点
	leaf.parent = r
	//当做根节点
	bt[-1] = r
	bt[0] = leaf


	return &bt

}

//返回根节点
func (bpt *BPlustree)Root()node{

	return (*bpt)[-1]

}

//处理第一个节点
func (bpt *BPlustree)First()node{

	return (*bpt)[0]

}

//统计数量
func (bpt *BPlustree)Count()int{
	count := 0
	leaf :=(*bpt)[0].(*LeafNode)
	for {
		//数量叠加
		count += leaf.CountNum()
		if leaf.next == nil{
			break
		}
		leaf = leaf.next
	}
	return count
}

func (bpt *BPlustree)Values()[]*LeafNode{
	nodes := make([]*LeafNode, 0)
	leaf := (*bpt)[0].(*LeafNode)
	for {
		//数据节点叠加
		nodes = append(nodes, leaf)
		if leaf.next == nil{
			break
		}
		leaf = leaf.next
	}
	return nodes
}

func (bpt *BPlustree)Insert(key int, value string){

}

func (bpt *BPlustree)Search(key int)(string,bool){
	//查找
	kv, _,_ = search((*bpt)[-1],key)
	if kv == nil{
		return "", false
	}else{
		return kv.value, true
	}
}

//搜索数据
func search(n node, key int)(*kv, int, *LeafNode){
	curr := n
	oldindex := -1
	for true {
		switch t := curr.(type) {
		case *LeafNode:
			i, ok := t.find(key)
			if !ok{
				return nil, oldindex,t

			}else{
				return &t.kvs[i],oldindex,t
			}
		case *interiorNode:
			i,_:=t.find(key)
			curr=t.kcs[i].child
			oldindex = i

		default:
			panic("error")
		}
	}
	return nil, 0, nil
}