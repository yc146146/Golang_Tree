package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Start int64
	End   int64
	Index int
	Max   int64
}

type Nodes []Node

func (ns *Nodes) Add(start, end int64) int {
	idx := len(*ns)
	*ns = append(*ns, Node{Start: start, Index: idx, End: end})
	//追加节点
	return idx
}

func max(a, b int64) int64 {
	if a < b {
		return b

	} else {
		return a

	}
}

//实现排序
func (ns Nodes) sort() {
	sort.Slice(ns, func(i, j int) bool {
		a, b := ns[i].Start, ns[j].Start
		switch {
		case a < b:
			return true
		case a > b:
			return false
		default:
			a, b = ns[i].End, ns[i].End
			return a < b
		}
	})
}

//递归寻找最大值
func (ns Nodes) FillMax(off int) int64 {
	length := len(ns)
	mid := length / 2
	v := ns[mid].End
	if mid > 0 {
		v = max(v, ns[:mid].FillMax(off))
	}

	if mid < length-1 {
		v = max(v, ns[mid+1:].FillMax(off))
	}

	ns[mid].Max = v

	return v

}



func (ns *Nodes) Build() {
	ns.sort()
	ns.FillMax(0)
}

func (ns Nodes) Query(q int64) []int {
	return ns.query(q, nil)
}

func (ns Nodes) query(q int64, res []int) []int {
	length := len(ns)
	mid := length / 2
	if q > ns[mid].Max {
		return res

	}
	if q >= ns[mid].Start && q <= ns[mid].End {
		res = append(res, ns[mid].Index)
	}
	if mid > 0 {
		res = ns[:mid].query(q, res)
	}

	if mid < length-1 {
		res = ns[mid+1:].query(q, res)
	}

	return res

}

func (ns Nodes) querynode(q int64, res []Node) []Node {
	length := len(ns)
	mid := length / 2
	if q > ns[mid].Max {
		return res

	}
	if q >= ns[mid].Start && q <= ns[mid].End {
		res = append(res, ns[mid])
	}
	if mid > 0 {
		res = ns[:mid].querynode(q, res)
	}

	if mid < length-1 {
		res = ns[mid+1:].querynode(q, res)
	}

	return res
}

func main() {
	var nn Nodes
	nn.Add(16, 21)
	nn.Add(8, 9)
	nn.Add(25, 30)
	nn.Add(5, 8)
	nn.Add(15, 23)
	nn.Add(17, 19)
	nn.Add(20, 26)
	nn.Add(0, 3)
	nn.Add(6, 10)
	nn.Add(19, 20)
	nn.Build()
	for i, x := range nn {
		fmt.Println("", i, x.Start, x.End, x.Max, x.Index)
	}
	r := nn.Query(18)

	fmt.Println("r",r)
}
