package main

import "fmt"

type Integer int
type Integers []int

func (i Integers) Merge(i2 MergerAble) MergerAble {

	newi2 := i2.(Integers)

	for j:=0;j<len(i);j++{
		i[j]+=newi2[j]
	}
	return i

}

func (i Integer) Merge(i2 MergerAble) MergerAble {
	return Integer(int(i) + int(i2.(Integer)))
}

//Comparable 可比较的
type Comparable interface {
	Compare(c2 Comparable) int
}

func (i Integer) Compare(c2 Comparable) int {
	return int(i) - int(c2.(Integer))
}

//MergerAble 可合并的
type MergerAble interface {
	Merge(m2 MergerAble) MergerAble
}

type ArraySegmentTree struct {
	data []MergerAble
	tree []MergerAble
	//merger Merger
}

//合并接口
func TransInts(original []int) []MergerAble {
	ret := make([]MergerAble, len(original))
	for key, value := range original {
		ret[key] = Integer(value)
	}

	return ret

}

//线段树
func CreateSegmentTree(arr []MergerAble) *ArraySegmentTree {
	//开辟4倍内存
	tree := make([]MergerAble, len(arr)*4)
	st := &ArraySegmentTree{arr, tree}
	st.BuildSegmentTree(0,0,len(arr)-1)
	return st

}

//取出左边
func (ast *ArraySegmentTree) LeftChild(index int) int {
	return index*2 + 1
}

//取出右边
func (ast *ArraySegmentTree) RightChild(index int) int {
	return index*2 + 2
}

//查看字符串
func (ast *ArraySegmentTree) String() string {
	return fmt.Sprintln(ast.tree)
}

//查看大小
func (ast *ArraySegmentTree) Size() int {
	return len(ast.data)
}

//构造树
func (ast *ArraySegmentTree) BuildSegmentTree(index, left, right int) {
	if left == right {
		//插入第一个元素
		ast.tree[index] = ast.data[left]
	} else {
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		//取得中间数据
		midchild := (right + left) / 2
		//反复构造
		ast.BuildSegmentTree(leftchild, left, midchild)
		ast.BuildSegmentTree(rightchild, midchild+1, right)
		//合并数据
		ast.tree[index] = ast.tree[leftchild].Merge(ast.tree[rightchild])
	}
}

//查询
func (ast *ArraySegmentTree) Query(qleft, qright int) MergerAble {
	if qleft<0 || qright<0 || qleft > len(ast.data) || qright>len(ast.data){
		panic("index is out")
	}

	return ast.query(0,0,len(ast.data)-1,qleft, qright)
}

func (ast *ArraySegmentTree) query(index, left, right, qleft, qright int) MergerAble {
	if left == right{
		//返回 找到
		return ast.tree[index]
	}else{
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		//取得中间数据
		midchild := (right + left) / 2

		if qleft >= midchild+1 {
			return ast.query(rightchild, midchild+1,right,qleft,qright)

		} else {
			return ast.query(leftchild, left,midchild,qleft,qright)
		}

		leftres := ast.query(leftchild, left, midchild, qleft,qright)
		rightres := ast.query(rightchild, midchild + 1, right, qleft,qright)
		return leftres.Merge(rightres)
	}
}


//设置
func (ast *ArraySegmentTree) Set(index int, e MergerAble) {
	if index < 0 || index >= len(ast.data){
		panic("index is out")
	}else{
		ast.data[index] = e
		ast.set(0,0, len(ast.data)-1, index, e)
	}
}

func (ast *ArraySegmentTree) set(tree, left, right, index int, e MergerAble) {
	if left == right {
		//左右平衡
		ast.tree[tree] = e

	} else {
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		//取得中间数据
		midchild := (right + left) / 2

		if index >= midchild+1 {
			ast.set(rightchild, midchild+1, right, index, e)
		} else {
			ast.set(leftchild, left, midchild, index, e)
		}
		//归并
		ast.tree[tree] = ast.tree[left].Merge(ast.tree[right])
	}
}

func main() {
	data := []MergerAble{Integer(1),Integer(2),Integer(13),Integer(23),Integer(43)}
	fmt.Println(data)
	mytree:=CreateSegmentTree(data)

	fmt.Println(mytree.data)
	fmt.Println(mytree.tree)

	mytree.Set(0, Integer(999))

	fmt.Println(mytree.data)
	fmt.Println(mytree.tree)

	fmt.Println(mytree.Query(2,4))
}