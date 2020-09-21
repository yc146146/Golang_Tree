package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	//正无穷大
	Inf = int(^uint(0) >> 1)
	//负无穷大
	NegInf = -Inf - 1
)

//线段的左右
type segment struct{
	from int
	to int

}

//区间
type interval struct{
	segment
	element interface{}
}

type node struct {
	//线段
	segment segment
	//二叉树的左右孩子
	left, right *node
	//指向指针
	intervals []*interval
}

type Tree struct {
	base []interval
	//元素索引
	elements map[interface{}]struct{}
	//根节点
	root *node
}


func (t *Tree)Push(from, to int, element interface{}){
	if to < from{
		//对角
		from, to = to, from

	}
	if t.elements==nil{
		//开辟内存
		t.elements=make(map[interface{}]struct{})
	}
	t.elements[element] = struct{}{}

	t.base = append(t.base, interval{segment{from,to}, element})
}

func (t *Tree)Clear(){
	t.base = nil
	t.root = nil

}

func (t *Tree)Buildtree()error{
	if len(t.base) == 0{
		return errors.New("已经构造")
	}else{
		//插入数据
		leaves := elementaryIntervals(t.endpoints())
		t.root=t.InsertNodes(leaves)
		for i:=range t.base{
			t.root.InsertInterval(&t.base[i])
		}

		return nil
	}


}


func removedups(sorted[]int)(unqique []int){
	unqique = make([]int, 0, len(sorted))
	unqique = append(unqique, sorted[0])
	prev := sorted[0]
	for _, val := range sorted[1:]{
		if val != prev{
			unqique = append(unqique, val)
			prev = val
		}
	}
	return
}

func (t *Tree)endpoints()[]int{
	baselen := len(t.base)
	endpoints := make([]int, baselen*2)
	for i,interval := range t.base{
		endpoints[i] = interval.from
		endpoints[i+baselen]=interval.to

	}

	sort.Sort(sort.IntSlice(endpoints))
	return removedups(endpoints)
}

func (t *Tree)InsertNodes(leaves[]segment)*node{
	var n*node
	if len(leaves)==1{
		n=&node{segment: leaves[0]}
		n.left=nil
		n.right=nil

	}else{
		n=&node{segment: segment{leaves[0].from, leaves[0].to}}
		//取得中间数据
		center := len(leaves)/2
		n.left = t.InsertNodes(leaves[:center])
		n.right = t.InsertNodes(leaves[center:])
	}

	return n

}



//区间判断
func (s*segment)subsetOf(other *segment)bool{
	return other.from <= s.from && other.to >= s.to
}

//区间判断
//要么other 包含s  要么 s包含ohter
func (s*segment)intersectswitch(other *segment)bool{
	return other.from <= s.from && other.to >= s.to ||  other.from >= s.from && other.to <= s.to
}

func (n *node)InsertInterval(i* interval){
	if n.segment.subsetOf(&i.segment){
		if n.intervals == nil{
			//开辟内存
			n.intervals = make([]*interval, 0,1)
		}
		//只有一个
		n.intervals=append(n.intervals, i)
	}else{
		if n.left!=nil && n.left.segment.intersectswitch(&i.segment){
			n.left.InsertInterval(i)
		}

		if n.right!=nil && n.right.segment.intersectswitch(&i.segment){
			n.right.InsertInterval(i)
		}
	}
}


func (t *Tree)Queryindex(index int){

}

func (s*segment)contains(index int)bool{
	return s.from<=index && s.to>=index

}

func (t*Tree)QueryIndex(index int)(<-chan interface{},error){
	if t.root==nil{
		return nil, errors.New("为空")
	}
	//构造管道
	intervals:=make(chan *interval)

	//并发调用
	go func(t*Tree, index int, intervals chan* interval) {
		query(t.root, index, intervals)
		//关闭管道
		close(intervals)
	}(t, index, intervals)

	elements := make(chan interface{})

	go func(intervals chan* interval, elements chan interface{}) {
		defer close(elements)
		results := make(map[interface{}]struct{})
		for interval:=range intervals{
			//找到
			_, alreadyFound := results[interval.element]
			if !alreadyFound{
				results[interval.element] = struct{}{}
				elements<-interval.element
				if len(results)>=len(t.elements){
					return
				}
			}
		}
	}(intervals,elements)

	return elements, nil

}

//查询数据
func query(node *node, index int, results chan <- *interval){
	//判断数据是否在区间内
	if node.segment.contains(index){
		for _, intervalx := range node.intervals{
			results <- intervalx
		}
		if node.left != nil{
			query(node.left, index, results)
		}

		if node.right != nil{
			query(node.right, index, results)
		}

	}
}

//区间划分
func elementaryIntervals(endpoints []int)[]segment{
	if len(endpoints) == 1{
		return []segment{{endpoints[0],endpoints[0]}}
	}else{
		intervals := make([]segment, len(endpoints)*2-1)
		for i:=0;i<len(endpoints);i++{
			intervals[i*2] = segment{endpoints[0], endpoints[0]}
			if i<len(endpoints)-1{
				intervals[2*i+1]=segment{endpoints[i],endpoints[i]}

			}
		}
		return intervals
	}
}

//遍历每一个节点
func Traverse(node *node, depth int, enter, leave func(*node, int)){
	if node == nil{
		return

	}
	if enter != nil{
		//递归函数
		enter(node, depth)
	}

	Traverse(node.left, depth+1, enter, leave)
	Traverse(node.right, depth+1, enter, leave)

	if leave!=nil{
		leave(node,depth)
	}

}

func log2(num int)int  {
	if num == 0{
		return NegInf
	}

	i:=-1
	for num>0{
		num=num>>1
		i++
	}
	return i

}

func space(n int){
	for i:=0;i<n;i++{
		//打印空格显示层级
		fmt.Print(" ")
	}
}

func (n *node)print(){
	//fmt.Println("n",n)
	from := fmt.Sprintf("%d", n.segment.from)
	switch n.segment.from {
	case Inf:
		from = "+00"
	case NegInf:
		from="-00"
	}

	to := fmt.Sprintf("%d", n.segment.to)
	switch n.segment.to {
	case Inf:
		from = "+00"
	case NegInf:
		from="-00"
	}

	//打印数据
	fmt.Printf("%s,%s", from, to)
	//fmt.Println("%v", n.intervals)
	if n.intervals != nil{
		fmt.Print("->[")

		for _,intervl := range n.intervals{
			fmt.Printf("(%v,%v)=[%v]", intervl.from, intervl.to, intervl.element)
		}

		fmt.Print("]")
	}
}

func (t*Tree)Print(){
	endpoints := len(t.base)*2+2
	leaves := endpoints*2-3
	//高度
	height := 1+log2(leaves)
	fmt.Println("height", height,"leaves", leaves)
	levels := make([][]*node, height+1)
	Traverse(t.root, 0, func(n *node, depth int) {
		levels[depth]=append(levels[depth],n)
	},nil)

	for i,level := range levels{
		for j,n := range level{
			space(12 * (len(levels)-1-i))
			n.print()
			space(1*(height-i))
			if j-1%2==0{
				space(2)
			}
		}
		fmt.Println()
	}
}


func main() {
	mytree := new(Tree)
	mytree.Push(1,10,"abcdefg")
	err:=mytree.Buildtree()
	fmt.Println(err)
	fmt.Println(mytree.QueryIndex(4))

	mytree.Print()
}